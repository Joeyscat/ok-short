package service

import (
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/labstack/echo/v4"
)

type GetLinkTraceRequest struct {
	Sc string `json:"sc" form:"sc" binding:"required"`
}

func (svc *Service) CreateLinkTrace(sc, url string, c echo.Context) error {
	ip := c.RealIP()
	ua := c.Request().UserAgent()
	cookies := c.Request().Cookies()
	var cookieStr string
	for _, cookie := range cookies {
		cookieStr += cookie.Name + ":" + cookie.Value + "&"
	}
	if len(cookieStr) > 2 {
		cookieStr = cookieStr[:len(cookieStr)-2]
	}
	lt := &model.LinkTrace{
		Sc:     sc,
		URL:    url,
		Ip:     ip,
		UA:     ua,
		Cookie: cookieStr,
	}

	return model.CreateLinkTrace(lt)
}

// 获取某短链访问记录，返回一个LinkTrace数组
func (svc *Service) GetLinkTrace(param *GetLinkTraceRequest, pager *app.Pager) ([]*model.LinkTrace, error) {
	return model.GetLinkTraceListBySc(param.Sc, int64(pager.Page), int64(pager.PageSize))
}

// 获取多个短链访问记录
func (svc *Service) GetLinkTraceList(pager *app.Pager) ([]*model.LinkTrace, error) {
	return model.GetLinkTraceList(int64(pager.Page), int64(pager.PageSize))
}
