package http

import (
	"bytes"
	"my-demo-service/c"
	"my-demo-service/config"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"gitlab.xxx.com/xxx-xxx/go-kit/logger"
	"go.elastic.co/apm"
)

func timeItMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spanQuery, _ := apm.StartSpan(ctx.Request.Context(), "queryxxx", "query")
		defer spanQuery.End()

		start := time.Now()

		// init header:
		allHeaders := initHeaders(ctx)

		// parse and reset body:
		var data []byte
		if ctx.Request.Body != nil {
			data, _ = ctx.GetRawData()
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
		spanQuery.Context.SetLabel("queryBody", string(data))

		ctx.Next()

		var panicTrace interface{}
		panicInfo, ok := ctx.Get("panic_trace")
		if ok {
			panicTrace = panicInfo
		}

		logger.LoadExtra(map[string]interface{}{
			"log_key":         "timeit_log",
			"service":         config.GetConfig().Service.Name,
			"user_agent":      ctx.Request.UserAgent(),
			"host":            ctx.Request.Host,
			"method":          ctx.Request.Method,
			"client_ip":       ctx.ClientIP(),
			"remote_addr":     ctx.Request.RemoteAddr,
			"body":            string(data),
			"request_headers": allHeaders,
			"search_panic":    panicTrace,
			"http_url":        ctx.Request.URL,
			"uri":             ctx.Request.RequestURI,
			"spent_time_ms":   time.Since(start).Milliseconds(),
			"spent_time_μs":   time.Since(start).Microseconds(),
		}).Info("xxx_timeit_log")
	}
}

// initHeaders return the headers per request
func initHeaders(ctx *gin.Context) map[string]string {
	headers := make(map[string]string)
	headers[c.Lang] = strings.ToLower(ctx.GetHeader(c.Lang))
	// request id:
	headers[c.XRequestID] = ctx.Request.Header.Get(c.XRequestID)
	// request id:
	requestID := ctx.Request.Header.Get(c.XRequestID)
	if requestID == "" {
		requestID = ksuid.New().String()
		ctx.Request.Header.Set(c.XRequestID, requestID)
	}
	headers[c.XRequestID] = requestID
	// vendor:
	vendor := strings.ToLower(ctx.GetHeader(c.Vendor))
	headers[c.Vendor] = vendor

	headers[c.CustomerID] = strings.ToUpper(ctx.GetHeader(c.HTTPHeaderCustomerID))
	headers[c.MediaCode] = strings.ToLower(ctx.GetHeader(c.HTTPHeaderMediaCode))
	return headers
}
