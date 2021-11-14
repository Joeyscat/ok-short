package v1

import (
	"github.com/joeyscat/ok-short/internal/global"
	srvv1 "github.com/joeyscat/ok-short/internal/service/v1"
	"github.com/joeyscat/ok-short/internal/store"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

type LinkTraceController struct {
	srv srvv1.Service
}

// NewLinkTraceController creates a link handler.
func NewLinkTraceController(store store.Factory) *LinkTraceController {
	return &LinkTraceController{
		srv: srvv1.NewService(store),
	}
}

func (ctl *LinkTraceController) Get(ctx echo.Context) error {
	sc := ctx.Param("sc")
	param := srvv1.GetLinkTraceRequest{Sc: sc}
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf(ctx.Request().Context(), "app.BindAndValid errs: %v", errs)
		return app.ErrorResponse(ctx, errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	traces, err := ctl.srv.LinkTraces().GetLinkTrace(ctx.Request().Context(), &param, &pager)
	if err != nil {
		global.Logger.Errorf(ctx.Request().Context(), "svc.GetLinkTrace err: %v", err)
		return app.ErrorResponse(ctx, errcode.ErrorGetLinkTraceFail)
	}

	return app.Response(ctx, map[string]interface{}{"traces": traces, "code": 0})
}

func (ctl *LinkTraceController) List(ctx echo.Context) error {
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	link, err := ctl.srv.LinkTraces().ListLinkTrace(ctx.Request().Context(), &pager)
	if err != nil {
		global.Logger.Errorf(ctx.Request().Context(), "svc.GetLinkList err: %v", err)
		return app.ErrorResponse(ctx, errcode.ErrorGetLinkListFail)
	}

	return app.Response(ctx, map[string]interface{}{"link": link, "code": 0})
}
