package controllers

import (
	"strconv"
	"strings"

	"my-demo-service/c"
	"my-demo-service/models"

	"my-demo-service/utils"

	"github.com/gin-gonic/gin"
	"gitlab.xxx.com/xxx-xxx/go-kit/logger"
)

func getRequestInfo(ctx *gin.Context) models.RequestInfo {
	var requestInfo models.RequestInfo
	// customerID
	requestInfo.CustomerID = strings.ToLower(getRequestData(ctx, c.HTTPHeaderCustomerID, c.CustomerID, c.GID))
	// userID
	requestInfo.UserID = ctx.Request.URL.Query().Get(c.UserID)
	//requestInfo.UserID = getRequestData(ctx, c.HTTPHeaderUserID, c.UserID, c.UserID)

	// requestID
	requestInfo.RequestID = ctx.Request.Header.Get(c.XRequestID)
	// vendor
	requestInfo.Vendor = ctx.Request.Header.Get(c.Vendor)
	// lang
	lang := ctx.Request.Header.Get(c.Lang)
	if lang == "" {
		lang = c.LangCn
	}
	requestInfo.Lang = lang
	return requestInfo
}

// getRequestData getRequestData
func getRequestData(ctx *gin.Context, headerKey string, urlKey string, otherKeys ...string) (value string) {
	if value = ctx.GetHeader(headerKey); value == "" {
		if value = ctx.Request.URL.Query().Get(urlKey); value == "" {
			for _, key := range otherKeys {
				if value = ctx.GetHeader(key); value != "" {
					break
				}
				if value = ctx.Request.URL.Query().Get(key); value != "" {
					break
				}
			}
		}
	}
	return value
}

// getIntFromURL getIntFromURL
func getIntFromURL(ctx *gin.Context, key string) (int, error) {
	campaignID := ctx.Request.URL.Query().Get(strings.ToLower(key))
	return strconv.Atoi(campaignID)
}

//getRealUID 获取真正UID  < 有且只有一个
func getRealUID(grayUID string) (uid string, isGrayLogin bool) {
	isGrayLogin = false
	grayUID = strings.TrimSpace(grayUID)
	if grayUID != "" {
		if index := strings.Index(grayUID, "<"); index > 0 {
			grayUID = grayUID[0:index]
			isGrayLogin = true
		}
	}
	return grayUID, isGrayLogin
}

// getIntArrayFromURL getIntArrayFromURL
func getIntArrayFromURL(ctx *gin.Context, key string) []int {
	var res []int
	value := ctx.Request.URL.Query().Get(strings.ToLower(key))
	values := strings.Split(value, ",")
	for _, id := range values {
		if temp, e := utils.Int(id); e == nil {
			res = append(res, temp)
		} else {
			logger.With("log_key", key).Error("params is invalid")
		}
	}
	return res
}

func formatResponse(resp *models.CommonResponse) *models.CommonResponse {
	if len(resp.Errors) > 0 {
		resp.Data = nil
		resp.Message = "wrong"
	} else {
		resp.Message = "ok"
	}

	return resp
}
