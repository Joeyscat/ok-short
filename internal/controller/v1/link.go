package v1

import (
	"net/http"

	"github.com/joeyscat/ok-short/internal/global"
	srvv1 "github.com/joeyscat/ok-short/internal/service/v1"
	"github.com/joeyscat/ok-short/internal/store"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

type LinkController struct {
	srv srvv1.Service
}

// NewLinkController creates a link handler.
func NewLinkController(store store.Factory) *LinkController {
	return &LinkController{
		srv: srvv1.NewService(store),
	}
}

func (ctl *LinkController) Shorten(ctx echo.Context) error {
	param := srvv1.CreateLinkRequest{}
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf(ctx.Request().Context(), "app.BindAndValid errs: %v", errs)
		return app.ErrorResponse(ctx, errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	link, err := ctl.srv.Links().CreateLink(ctx.Request().Context(), &param)
	if err != nil {
		global.Logger.Errorf(ctx.Request().Context(), "svc.CreateLink err: %v", err)
		return app.ErrorResponse(ctx, errcode.ErrorCreateLinkFail)
	}

	return app.Response(ctx, map[string]interface{}{"link": link, "code": errcode.Success.Code})
}

func (ctl *LinkController) Redirect(ctx echo.Context) error {
	sc := ctx.Param("sc")
	param := srvv1.RedirectLinkRequest{Sc: sc}
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf(ctx.Request().Context(), "app.BindAndValid errs: %v", errs)
		return app.ErrorResponse(ctx, errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	link, err := ctl.srv.Links().UnShorten(ctx.Request().Context(), &param)
	if err != nil {
		global.Logger.Errorf(ctx.Request().Context(), "svc.UnShorten err: %v", err)
		return app.ErrorResponse(ctx, errcode.ErrorUnShortLinkFail)
	}
	if link == "" {
		return ctx.Redirect(http.StatusTemporaryRedirect, "http://oook.fun/404")
	}

	// save link trace
	go func() {
		err = ctl.srv.LinkTraces().CreateLinkTrace(ctx.Request().Context(), sc, link, ctx)
		if err != nil {
			global.Logger.Errorf(ctx.Request().Context(), "svc.CreateLinkTrace err: %v", err)
		}
	}()

	return ctx.Redirect(http.StatusTemporaryRedirect, link)
}

func (ctl *LinkController) Get(ctx echo.Context) error {
	sc := ctx.Param("sc")
	param := srvv1.GetLinkRequest{Sc: sc}
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf(ctx.Request().Context(), "app.BindAndValid errs: %v", errs)
		return app.ErrorResponse(ctx, errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	link, err := ctl.srv.Links().GetLink(ctx.Request().Context(), &param)
	if err != nil {
		global.Logger.Errorf(ctx.Request().Context(), "svc.GetLink err: %v", err)
		return app.ErrorResponse(ctx, errcode.ErrorGetLinkFail)
	}

	return app.Response(ctx, map[string]interface{}{"link": link, "code": errcode.Success.Code})
}

func (c *LinkController) List(ctx echo.Context) error {
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	link, err := c.srv.Links().ListLinks(ctx.Request().Context(), &pager)
	if err != nil {
		global.Logger.Errorf(ctx.Request().Context(), "svc.GetLinkList err: %v", err)
		return app.ErrorResponse(ctx, errcode.ErrorGetLinkListFail)
	}

	return app.Response(ctx, map[string]interface{}{"link": link, "code": errcode.Success.Code})
}
