package service

import (
	"fmt"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/msg"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/codec"
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

	// 存储短链接代码与原链接的映射到Redis SC:URL
	expiration := time.Minute * time.Duration(param.ExpirationInMinutes)
	err = global.Redis.Set(fmt.Sprintf(LinkKey, sc), param.URL, expiration).Err()
	if err != nil {
		return "", err
	}
	// TODO 定时清理过期数据
	// 存储短链详情到数据库
	linkMsg := &msg.LinkMsg{Sc: sc, URL: param.URL, Exp: param.ExpirationInMinutes}
	msgBytes, err := codec.Encoder(linkMsg)
	err = global.Nats.Publish(global.NatsSetting.Subj.LinkDetail, msgBytes)

	return url, err
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
