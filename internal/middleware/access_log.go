package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mneumi/gin-boilerplate/global"
	"github.com/mneumi/gin-boilerplate/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		// 替换原有 ctx.Writer，实现双写
		ctx.Writer = bodyWriter

		beginTime := time.Now().Unix()
		ctx.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		global.Logger.WithFields(fields).
			Infof("access log: method %s, status_code: %d, begin_time: %d, end_time: %d",
				ctx.Request.Method, bodyWriter.Status(), beginTime, endTime)
	}
}
