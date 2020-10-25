package service

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/internel/model"
	"github.com/joeyscat/ok-short/pkg/app"
)

type GetLinkTraceRequest struct {
	Sc string `json:"sc" form:"sc" binding:"required"`
}

func (svc *Service) CreateLinkTrace(sc, url string, c *gin.Context) (*model.LinkTrace, error) {
	ip := c.ClientIP()
	ua := c.Request.UserAgent()
	cookies := c.Request.Cookies()
	var cookieStr string
	for _, cookie := range cookies {
		cookieStr += cookie.Name + ":" + cookie.Value + "&"
	}

	return svc.dao.CreateLinkTrace(sc, url, ip, ua, cookieStr)
}

// 获取某短链访问记录，返回一个LinkTrace数组
func (svc *Service) GetLinkTrace(param *GetLinkTraceRequest, pager *app.Pager) ([]*model.LinkTrace, error) {
	return svc.dao.GetLinkTrace(param.Sc, pager.Page, pager.PageSize)
}

// 获取多个短链访问记录
func (svc *Service) GetLinkTraceList(pager *app.Pager) ([]*model.LinkTrace, error) {
	// TODO
	return nil, nil
}
