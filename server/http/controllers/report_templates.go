package controllers

import (
	"my-demo-service/c"
	"my-demo-service/models"
	xxxtemplates "my-demo-service/service/xxx_templates"
	"my-demo-service/utils"
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

// GetRTList template api
func GetRTList(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxTemplate
	reqInfo := getRequestInfo(ctx)
	data, errs := xxxtemplates.NewxxxTemplates().GetList(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	resp.Data = data
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}

// GetRTOne template api
func GetRTOne(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxTemplate
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
	data, errs := xxxtemplates.NewxxxTemplates().GetOne(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	resp.Data = data
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}

// RTCreate template api
func RTCreate(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxTemplate
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
	if !utils.CheckAnalysis(req.AnalysisType) {
		resp.Errors.AddError(errors.New("analysis type wrong"), "500")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if utf8.RuneCountInString(req.Name) > 30 {
		resp.Errors.AddError(errors.New("name length beyond 30 character"), "400")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	reqInfo := getRequestInfo(ctx)
	data, errs := xxxtemplates.NewxxxTemplates().Create(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	resp.Data = data
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}

// RTUpdate template api
func RTUpdate(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxTemplate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Errors.AddError(err, "500")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if !utils.CheckAnalysis(req.AnalysisType) {
		resp.Errors.AddError(errors.New("analysis type wrong"), "500")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if utf8.RuneCountInString(req.Name) > 30 {
		resp.Errors.AddError(errors.New("name length beyond 30 character"), "400")
		ctx.JSON(http.StatusOK, resp)
		return
	}
	reqInfo := getRequestInfo(ctx)
	data, errs := xxxtemplates.NewxxxTemplates().Update(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	resp.Data = data
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}

// RTDelete template api
func RTDelete(ctx *gin.Context) {
	var resp models.CommonResponse
	var req models.xxxTemplate
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
	_, errs := xxxtemplates.NewxxxTemplates().Delete(ctx.Request.Context(), reqInfo, req)
	resp.Errors = append(resp.Errors, errs...)
	formatResponse(&resp)
	ctx.JSON(http.StatusOK, resp)
}
