package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mneumi/gin-boilerplate/global"
	"github.com/mneumi/gin-boilerplate/pkg/app"
	"github.com/mneumi/gin-boilerplate/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
