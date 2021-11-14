package middleware

import (
	"time"

	"github.com/joeyscat/ok-short/internal/global"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

// 请求限制：
// 1.IP限流(未登录的用户请求)
// 2.token作为用户标识进行限流
func RequestLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := "req:times:" + c.RealIP()
		_, err := global.Redis.Get(key).Result()

		if err != nil {
			global.Redis.Set(key, true, time.Millisecond*global.AppSetting.RequestLimit)
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}

		return app.ErrorResponse(c, errcode.TooManyRequests)
	}
}
