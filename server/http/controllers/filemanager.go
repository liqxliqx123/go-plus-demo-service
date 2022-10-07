package controllers

import (
	"fmt"
	"net/http"

	"my-demo-service/app"
	"my-demo-service/c"
	"my-demo-service/models"
	"my-demo-service/utils"

	"github.com/gin-gonic/gin"
)

// PreSinged controller
func PreSinged(ctx *gin.Context) {
	var resp models.CommonResponse
	var payload models.PreSignedPayload

	span, _ := utils.StartSpanForExported(ctx, "presigned")
	defer span.End()
	if err := ctx.ShouldBindQuery(&payload); err != nil || payload.Extension == "" {
		ctx.JSON(http.StatusBadRequest, models.CommonResponse{
			Errors: []models.Error{{
				Message: err.Error(),
				Code:    c.HTTPInvalidParamsCode,
			}},
			Message: c.ErrMessage,
		})
		return
	}
	application := app.Default()
	filename := fmt.Sprintf("logo%s", payload.Extension)
	uri, formData, err := application.FileManager.PreSigned(ctx, filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.CommonResponse{
			Errors: []models.Error{{
				Message: err.Error(),
				Code:    c.HTTPInvalidParamsCode,
			}},
			Message: c.ErrMessage,
		})
		return
	}
	resp.Data = models.PreSingedResponse{
		URL:      uri.String(),
		FormData: formData,
	}
	resp.Message = "ok"
	// s, _ := application.FileManager.GetFileURL(ctx, "common-filestore-dev", "xxx/c4d22944-29b5-11ed-a304-1e0031148645/img.jpg", "user1", "we6test1")
	// fmt.Println(s)
	ctx.JSON(http.StatusOK, resp)
}
