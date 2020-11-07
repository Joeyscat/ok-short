package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internel/service"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"net/http"
)

type Link struct{}

func NewLink() Link {
	return Link{}
}

// @Summary 新增短链接
// @Accept  json
// @Produce  json
// @Param link body service.CreateLinkRequest true "链接信息"
// @Success 200 {object} model.Link "成功，短链数据"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/links [post]
func (t Link) Shorten(c *gin.Context) {
	param := service.CreateLinkRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	link, err := svc.CreateLink(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateLink err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateLinkFail)
		return
	}

	response.ToResponse(gin.H{"link": link, "code": 0})
	return
}

// @Summary 短链接跳转
// @Produce  json
// @Param sc path string true "短链接ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /a/{sc} [get]
func (t Link) Redirect(c *gin.Context) {
	sc := c.Param("sc")
	param := service.RedirectLinkRequest{Sc: sc}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	link, err := svc.UnShorten(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UnShorten err: %v", err)
		response.ToErrorResponse(errcode.ErrorUnShortLinkFail)
		return
	}

	// save visit log
	svc.CreateLinkTrace(sc, link, c)

	c.Redirect(http.StatusTemporaryRedirect, link)
}

// @Summary 获取单个短链详情
// @Produce json
// @Param sc path string true "短链ID"
// @Success 200 {object} model.Link "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/links/{sc} [get]
func (t Link) Get(c *gin.Context) {
	sc := c.Param("sc")
	param := service.GetLinkRequest{Sc: sc}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	link, err := svc.GetLink(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetLink err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetLinkFail)
		return
	}

	response.ToResponse(gin.H{"link": link, "code": 0})
	return
}

// @Summary 获取多个短链接
// @Produce json
// @Success 200 {object} model.LinkSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/links [get]
func (t Link) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	link, err := svc.GetLinkList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetLinkList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetLinkListFail)
		return
	}

	response.ToResponse(gin.H{"link": link, "code": 0})
	return
}
