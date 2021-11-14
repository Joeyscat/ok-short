package v2

import (
	"net/http"

	"github.com/joeyscat/ok-short/internal/global"
	"github.com/joeyscat/ok-short/internal/service"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

type Link struct{}

func NewLink() Link {
	return Link{}
}

func (t Link) Shorten(e echo.Context) error {
	param := service.CreateLinkRequest{}
	response := app.NewResponse(e)
	valid, errs := app.BindAndValid(e, &param)
	if !valid {
		global.Logger.Errorf(e.Request().Context(), "app.BindAndValid errs: %v", errs)
		return response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(e.Request().Context())
	link, err := svc.CreateLink(&param)
	if err != nil {
		global.Logger.Errorf(e.Request().Context(), "svc.CreateLink err: %v", err)
		return response.ToErrorResponse(errcode.ErrorCreateLinkFail)
	}

	return response.ToResponse(map[string]interface{}{"link": link, "code": errcode.Success.Code})
}

func (t Link) Redirect(e echo.Context) error {
	sc := e.Param("sc")
	param := service.RedirectLinkRequest{Sc: sc}
	valid, errs := app.BindAndValid(e, &param)
	response := app.NewResponse(e)
	if !valid {
		global.Logger.Errorf(e.Request().Context(), "app.BindAndValid errs: %v", errs)
		return response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(e.Request().Context())
	link, err := svc.UnShorten(&param)
	if err != nil {
		global.Logger.Errorf(e.Request().Context(), "svc.UnShorten err: %v", err)
		return response.ToErrorResponse(errcode.ErrorUnShortLinkFail)
	}
	if link == "" {
		return e.Redirect(http.StatusTemporaryRedirect, "http://oook.fun/404")
	}

	// save link trace
	go func() {
		err = svc.CreateLinkTrace(sc, link, e)
		if err != nil {
			global.Logger.Errorf(e.Request().Context(), "svc.CreateLinkTrace err: %v", err)
		}
	}()

	return e.Redirect(http.StatusTemporaryRedirect, link)
}

func (t Link) Get(e echo.Context) error {
	sc := e.Param("sc")
	param := service.GetLinkRequest{Sc: sc}
	response := app.NewResponse(e)
	valid, errs := app.BindAndValid(e, &param)
	if !valid {
		global.Logger.Errorf(e.Request().Context(), "app.BindAndValid errs: %v", errs)
		return response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(e.Request().Context())
	link, err := svc.GetLink(&param)
	if err != nil {
		global.Logger.Errorf(e.Request().Context(), "svc.GetLink err: %v", err)
		return response.ToErrorResponse(errcode.ErrorGetLinkFail)
	}

	return response.ToResponse(map[string]interface{}{"link": link, "code": errcode.Success.Code})
}

func (t Link) List(e echo.Context) error {
	response := app.NewResponse(e)
	svc := service.New(e.Request().Context())

	pager := app.Pager{Page: app.GetPage(e), PageSize: app.GetPageSize(e)}
	link, err := svc.GetLinkList(&pager)
	if err != nil {
		global.Logger.Errorf(e.Request().Context(), "svc.GetLinkList err: %v", err)
		return response.ToErrorResponse(errcode.ErrorGetLinkListFail)
	}

	return response.ToResponse(map[string]interface{}{"link": link, "code": errcode.Success.Code})
}
