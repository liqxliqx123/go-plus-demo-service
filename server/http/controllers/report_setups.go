package controllers

import (
	"my-demo-service/c"
	"my-demo-service/models"
	xxxSetups "my-demo-service/service/xxx_setups"
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

// GetRSOne template api
func GetRSOne(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxSetup
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CommonResponse{
			Errors: []models.Error{{
				Message: err.Error(),
				Code:    c.HTTPInvalidParamsCode,
			}},
			Message: c.ErrMessage,
		})
		return
	}
	reqInfo := getRequestInfo(ctx)
	data, errs := xxxSetups.NewxxxSetups().GetOne(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	resp.Data = data
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}

// RSCreate template api
func RSCreate(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxSetup
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Errors.AddError(err, "500")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if req.Name == "" {
		resp.Errors.AddError(errors.New("name empty"), "500")
		ctx.JSON(http.StatusOK, resp)
		return
	}

	if utf8.RuneCountInString(req.Name) > 30 {
		resp.Errors.AddError(errors.New("name length beyond 30 character"), "400")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	reqInfo := getRequestInfo(ctx)
	data, errs := xxxSetups.NewxxxSetups().Create(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	resp.Data = data
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}

// RSUpdate template api
func RSUpdate(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxSetup
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Errors.AddError(err, "500")
		ctx.JSON(http.StatusOK, resp)
		return
	}

	if utf8.RuneCountInString(req.Name) > 30 {
		resp.Errors.AddError(errors.New("name length beyond 30 character"), "400")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	reqInfo := getRequestInfo(ctx)
	data, errs := xxxSetups.NewxxxSetups().Update(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	resp.Data = data
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}

// RSDelete template api
func RSDelete(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxSetup
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CommonResponse{
			Errors: []models.Error{{
				Message: err.Error(),
				Code:    c.HTTPInvalidParamsCode,
			}},
			Message: c.ErrMessage,
		})
		return
	}
	reqInfo := getRequestInfo(ctx)
	_, errs := xxxSetups.NewxxxSetups().Delete(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}
