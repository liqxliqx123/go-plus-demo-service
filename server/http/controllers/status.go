package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status status api
func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
