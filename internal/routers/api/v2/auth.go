package v2

import (
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/service"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

func GetAuth(e echo.Context) error {
	param := service.AuthRequest{}
	response := app.NewResponse(e)
	valid, errs := app.BindAndValid(e, &param)
	if !valid {
		global.Logger.Errorf(e.Request().Context(), "app.BindAndValid errs: %v", errs)
		return response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(e.Request().Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(e.Request().Context(), "svc.CheckAuth err: %v", err)
		return response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(e.Request().Context(), "app.GenerateToken err: %v", err)
		return response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	}

	return response.ToResponse(map[string]interface{}{
		"token": token,
	})
}
