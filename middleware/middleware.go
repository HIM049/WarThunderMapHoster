package middleware

import (
	"net/http"
	"thunder_hoster/config"
	"thunder_hoster/public"
	"time"

	"github.com/gin-gonic/gin"
)

// 访问控制
func AccessControlMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if public.ValidTime.Before(time.Now()) {
			// 已经超时
			ctx.AbortWithStatus(http.StatusServiceUnavailable)
		}
		ctx.Next()
	}
}

// 过多请求拒绝
func FailedCountLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if public.FailedCounter.Get(ctx.ClientIP()) > config.Cfg.RetryCount {
			ctx.AbortWithStatus(http.StatusTooManyRequests)
		}
		ctx.Next()
	}
}
