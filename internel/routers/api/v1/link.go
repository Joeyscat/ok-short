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
// @Param url path string true "长链接" minlength(10) maxlength(100)
// @Param expiration_in_minutes path int false "有效时间" min(0) max(1440) default(1440)
// @Success 200 {object} model.Link "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/links [post]
func (t Link) Shorten(c *gin.Context) {
	param := service.CreateLinkRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	link, err := svc.CreateLink(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateLink err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateLinkFail)
		return
	}

	response.ToResponse(gin.H{"link": link, "code": 0})
	return
}

// @Summary 短链接跳转
// @Produce  json
// @Param id path int true "短链接ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/links/{id} [delete]
func (t Link) Redirect(c *gin.Context) {
	param := service.RedirectLinkRequest{Sc: c.Param("sc")}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	link, err := svc.UnShorten(&param)
	if err != nil {
		global.Logger.Errorf("svc.UnShorten err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateLinkFail)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, link)
}
