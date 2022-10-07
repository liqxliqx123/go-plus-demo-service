package app

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"my-demo-service/config"
	"my-demo-service/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var defaultApplication *Application

// Application services
type Application struct {
	mu sync.Mutex

	Config     *config.Config
	Logger     *zap.SugaredLogger
	DBManager  db.DatabaseSourceManager
	HTTPServer *gin.Engine
	GRPCServer *grpc.Server
	GORMDB     *gorm.DB

	// Init and destroy hooks
	initHooks    []ApplicationHook
	destroyHooks []ApplicationHook
}

// ApplicationHook 应用启动时和小会时回调函数
type ApplicationHook func(*Application) error

// Default 获取默认的服务
func Default() *Application {
	if defaultApplication != nil {
		return defaultApplication
	}
	o := &Application{
		Config: config.GetConfig(),

		initHooks:    make([]ApplicationHook, 0),
		destroyHooks: make([]ApplicationHook, 0),
	}

	initLoggerApplicationHook(o)
	initGinApplicationHook(o)

	o.AddInitHook(initGRPCServerApplicationHook)
	// Uncomment next line if need database manager
	o.AddInitHook(initDatabasesApplicationHook)
	o.AddInitHook(initGormDatabasesApplicationHook)
	defaultApplication = o

	return o
}

// AddInitHook 添加初始化回调函数
func (app *Application) AddInitHook(f ApplicationHook) {
	app.initHooks = append(app.initHooks, f)
}

// AddDestroyHook 添加应用程序结束回调函数
func (app *Application) AddDestroyHook(f ApplicationHook) {
	app.destroyHooks = append(app.destroyHooks, f)
}

// Environment 获取环境
func (app *Application) Environment() string {
	return strings.ToLower(app.Config.Env)
}

// IsProduction 判断当前是否是生产环境
func (app *Application) IsProduction() bool {
	return app.Environment() == EnvProduction
}

// Run 启动服务
func (app *Application) Run() {
	app.callInitHooks()

	errc := make(chan error)

	go func() {
		addr := fmt.Sprintf("%s:%d", app.Config.Service.Host, app.Config.Service.Port)
		if app.Logger != nil {
			app.Logger.Info("running http/grpc server on: ", addr)
		}
		errc <- http.ListenAndServe(addr, app.HTTPHandler())
	}()

	panic(fmt.Sprintf("application run error: %s", <-errc))
}

// DryRun run the application but without http server
func (app *Application) DryRun() {
	app.callInitHooks()
}

// Shutdown 关闭服务
func (app *Application) Shutdown() {
	app.mu.Lock()
	defer app.mu.Unlock()

	if app.Logger != nil {
		app.Logger.Warn("shutdowning")
	}

	done := make(chan struct{})
	go app.callDestroyHooks(done)
	t := time.NewTimer(5 * time.Second)

	select {
	case <-done:
		break
	case <-t.C:
		if app.Logger != nil {
			app.Logger.Warn("timeout: application destroy hooks interrupted")
		}
		break
	}

	os.Exit(0)
}

func (app *Application) callInitHooks() {
	for _, hook := range app.initHooks {
		if err := hook(app); err != nil {
			panic(err)
		}
	}
}

func (app *Application) callDestroyHooks(done chan struct{}) {
	for i := len(app.destroyHooks); i > 0; i-- {
		hook := app.destroyHooks[i-1]
		if err := hook(app); err != nil {
			app.Logger.Error("calling application destroy hook error: ", err.Error())
		}
	}

	done <- struct{}{}
}

// HTTPHandler serve conn with GRPC or HTTP server
func (app *Application) HTTPHandler() http.Handler {
	return h2c.NewHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			contentType := r.Header.Get("Content-Type")

			for i, c := range contentType {
				if c == ' ' || c == ';' {
					contentType = contentType[:i]
					break
				}
			}

			if r.ProtoMajor == 2 && contentType == "application/grpc" {
				app.GRPCServer.ServeHTTP(w, r)
			} else {
				app.HTTPServer.ServeHTTP(w, r)
			}
		}),
		&http2.Server{},
	)
}
