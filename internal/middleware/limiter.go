package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mneumi/gin-boilerplate/pkg/app"
	"github.com/mneumi/gin-boilerplate/pkg/errcode"
	"github.com/mneumi/gin-boilerplate/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
