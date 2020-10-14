package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internel/service"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
)

type LinkTrace struct{}

func NewLinkTrace() LinkTrace {
	return LinkTrace{}
}

// @Summary 获取短链的访问记录
// @Produce json
// @Param id path int true "短链sc"
// @Success 200 {object} model.Link "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/link-trace/{sc} [get]
func (t LinkTrace) Get(c *gin.Context) {
	sc := c.Param("sc")
	param := service.GetLinkTraceRequest{Sc: sc}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	traces, err := svc.GetLinkTrace(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetLinkTrace err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetLinkTraceFail)
		return
	}

	response.ToResponse(gin.H{"traces": traces, "code": 0})
	return
}

// @Summary 获取多个短链接访问记录
// @Produce json
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.LinkSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/link-trace [get]
func (t LinkTrace) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	link, err := svc.GetLinkTraceList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetLinkList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetLinkListFail)
		return
	}

	response.ToResponse(gin.H{"link": link, "code": 0})
	return
}
