package v2

import (
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/service"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

type LinkTrace struct{}

func NewLinkTrace() LinkTrace {
	return LinkTrace{}
}

func (t LinkTrace) Get(c echo.Context) error {
	sc := c.Param("sc")
	param := service.GetLinkTraceRequest{Sc: sc}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c.Request().Context(), "app.BindAndValid errs: %v", errs)
		return response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(c.Request().Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	traces, err := svc.GetLinkTrace(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c.Request().Context(), "svc.GetLinkTrace err: %v", err)
		return response.ToErrorResponse(errcode.ErrorGetLinkTraceFail)
	}

	return response.ToResponse(map[string]interface{}{"traces": traces, "code": 0})
}

func (t LinkTrace) List(c echo.Context) error {
	response := app.NewResponse(c)
	svc := service.New(c.Request().Context())

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	link, err := svc.GetLinkTraceList(&pager)
	if err != nil {
		global.Logger.Errorf(c.Request().Context(), "svc.GetLinkList err: %v", err)
		return response.ToErrorResponse(errcode.ErrorGetLinkListFail)
	}

	return response.ToResponse(map[string]interface{}{"link": link, "code": 0})
}
