package service

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/internel/model"
)

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
