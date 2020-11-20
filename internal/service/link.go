package service

import (
	"fmt"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
	"time"
)

const (
	// URLIdKey redis自增主键所用的key
	URLIdKey = "next.url.id"
	// LinkKey 用于保存短链与原始链接的映射
	LinkKey = "link:%s:url"
)

type CreateLinkRequest struct {
	URL                 string `json:"url" form:"url" binding:"required,min=20,max=200"`
	ExpirationInMinutes uint32 `json:"expiration_in_minutes" form:"created_by" binding:"min=0,max=1440"`
}

type GetLinkRequest struct {
	Sc string `json:"sc" form:"sc" binding:"required"`
}

type RedirectLinkRequest struct {
	Sc string `json:"sc"  form:"sc" binding:"required,min=1,max=6"`
}

func (svc *Service) CreateLink(param *CreateLinkRequest) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := genId()
	sc := app.Base62Encode(id)
	url := global.AppSetting.LinkPrefix + sc

	expiration := time.Minute * time.Duration(param.ExpirationInMinutes)
	err = global.Redis.Set(fmt.Sprintf(LinkKey, sc), param.URL, expiration).Err()
	if err != nil {
		return "", err
	}
	// 存储原链接与短链接代码的映射
	// TODO 定时清理过期数据
	_, err = svc.dao.CreateLink(sc, param.URL, param.ExpirationInMinutes)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (svc *Service) UnShorten(param *RedirectLinkRequest) (string, error) {
	return global.Redis.Get(fmt.Sprintf(LinkKey, param.Sc)).Result()
	//link, err := svc.dao.GetLink(param.Sc)
	//if err != nil {
	//	return "", err
	//}
	//return link.OriginURL, nil
}

func (svc *Service) GetLink(param *GetLinkRequest) (*model.Link, error) {
	return svc.dao.GetLink(param.Sc)
}

func (svc *Service) GetLinkList(pager *app.Pager) ([]*model.Link, error) {
	return svc.dao.GetLinkList(pager.Page, pager.PageSize)
	// TODO DO -> VO
}

func genId() (int64, error) {
	// TODO should lock #1 begin
	// increase the global counter
	err := global.Redis.Incr(URLIdKey).Err()
	if err != nil {
		return -1, err
	}

	// encode global counter to base62
	return global.Redis.Get(URLIdKey).Int64()
	// TODO should lock #1 end
}
