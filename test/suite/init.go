package suite

import (
	"my-demo-service/app"
	"my-demo-service/config"
	httpServer "my-demo-service/server/http"
	"net"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	kitconfig "gitlab.xxx.com/xxx-xxx/go-kit/utils/config"
)

// SetupSuite run before all suites
func (suite *APISuite) SetupSuite() {
	suite.setupConfig()
	suite.setupServer()
}

func (suite *APISuite) setupConfig() {
	_, file, _, _ := runtime.Caller(0)
	configDir := filepath.Join(filepath.Dir(file), "../env/")

	kitconfig.LoadConf(configDir, config.GetConfig())
}

func (suite *APISuite) setupServer() {
	gin.SetMode(gin.TestMode)

	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		panic(err)
	}

	suite.Listener = lis

	server := app.Default()

	server.Config.Service.Port = lis.Addr().(*net.TCPAddr).Port

	server.DryRun()

	suite.App = server
	httpServer.RegisterRoutes(server.HTTPServer)
	go http.Serve(lis, suite.App.HTTPHandler())

	server.Logger.Info("running testing server at ", lis.Addr())
}
