package main

import (
	"my-demo-service/app"
	"my-demo-service/config"
	"my-demo-service/server/http"
	"os"
	"os/signal"
	"syscall"

	kitConfig "gitlab.xxx.com/xxx-xxx/go-kit/utils/config"

	_ "github.com/go-sql-driver/mysql"

	_ "go.uber.org/automaxprocs"
)

var cfgFile string

func init() {
	kitConfig.LoadConf(".", config.GetConfig())
}

func handleSignals(server *app.Application) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	server.Logger.Infof("signal %s received", <-sigs)
	server.Shutdown()
}

// @title my-demo-SERVICE API
// @version 1.0
// @description this is my-demo-service service api
// @contact.name Repository
// @contact.url https://gitlab.xxx.com/xxx-xxx/my-demo-service
// @host localhost:8080
// @query.collection.format multi
// @schemes http
func main() {
	server := app.Default()
	http.RegisterRoutes(server.HTTPServer)
	go handleSignals(server)
	server.Run()
}
