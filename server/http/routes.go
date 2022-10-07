package http

import (
	"my-demo-service/server/http/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(router gin.IRouter) {
	apiVersion := "v1"

	// status
	//router.Any("status", controllers.Status)

	// publicRouter DOES NOT ENFORCE the authentication
	publicRouter := router.Group(apiVersion, timeItMiddleware())

	// xxx-templates
	publicRouter.GET("xxx-templates", controllers.GetRTList)
	publicRouter.POST("xxx-templates", controllers.RTCreate)
	publicRouter.DELETE("xxx-templates/:xxx_template_id", controllers.RTDelete)
	publicRouter.PATCH("xxx-templates", controllers.RTUpdate)
	publicRouter.GET("xxx-templates/:xxx_template_id", controllers.GetRTOne)

	// xxx-setups
	publicRouter.POST("xxx-setups", controllers.RSCreate)
	publicRouter.DELETE("xxx-setups/:xxx_setup_id", controllers.RSDelete)
	publicRouter.PATCH("xxx-setups", controllers.RSUpdate)
	publicRouter.GET("xxx-setups/:xxx_setup_id", controllers.GetRSOne)
	// publicRouter.GET("xxx-setups", controllers.GetList)
	// minio
	toolRouter := publicRouter.Group("tool")
	// pre-signed
	toolRouter.GET("presigned", controllers.PreSinged)

}
