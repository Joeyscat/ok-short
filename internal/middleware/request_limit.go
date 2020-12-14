package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"time"
)

// 请求限制：
// 1.IP限流(未登录的用户请求)
// 2.token作为用户标识进行限流
func RequestLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "req:times:" + c.ClientIP()
		_, err := global.Redis.Get(key).Result()

		if err != nil {
			global.Redis.Set(key, true, time.Millisecond*global.AppSetting.RequestLimit)
			c.Next()
			return
		}

		app.NewResponse(c).ToErrorResponse(errcode.TooManyRequests)
		c.Abort()
	}
}
