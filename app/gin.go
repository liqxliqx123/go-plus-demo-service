package app

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
	"net/http"
)

func initGinApplicationHook(app *Application) error {
	if app.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.EnableJsonDecoderUseNumber()

	app.HTTPServer = gin.New()
	app.HTTPServer.Use(gin.Recovery())
	app.HTTPServer.Use(apmgin.Middleware(app.HTTPServer))

	app.Logger.Debug("init gin http server")
	app.HTTPServer.GET("/status", Status)

	return nil
}

// Status status api
func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
