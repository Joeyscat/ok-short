package middleware

import (
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

func Recovery(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().
					Errorf(c.Request().Context(), "panic recover err: %v", err)

				// TODO Email
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
			}
		}()
		return next(c)
	}
}
