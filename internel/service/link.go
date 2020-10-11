package service

import (
	"fmt"
	"github.com/joeyscat/ok-short/internel/pkg"
	"github.com/joeyscat/ok-short/pkg/app"
	"time"
)

type CreateLinkRequest struct {
	URL                 string `json:"url" form:"url" binding:"required,min=20,max=200"`
	ExpirationInMinutes uint32 `json:"expiration_in_minutes" form:"created_by" binding:"min=0,max=1440"`
}

type RedirectLinkRequest struct {
	Sc string `json:"sc"  form:"sc" binding:"required,min=1,max=6"`
}

func (svc *Service) CreateLink(param *CreateLinkRequest) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := genId()
	sc := app.Base62Encode(id)
	url := pkg.LinkPrefix + sc

	expiration := time.Minute * time.Duration(param.ExpirationInMinutes)
	err = pkg.ReCli.Set(fmt.Sprintf(pkg.LinkKey, sc), param.URL, expiration).Err()
	if err != nil {
		return "", err
	}
	// 存储原链接与短链接代码的映射
	// TODO 定时清理过期数据
	err = svc.dao.CreateLink(sc, param.URL, param.ExpirationInMinutes)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (svc *Service) UnShorten(param *RedirectLinkRequest) (string, error) {
	link := svc.dao.GetLink(param.Sc)
	return link.OriginURL, nil
}

func genId() (int64, error) {
	// TODO should lock #1 begin
	// increase the global counter
	err := pkg.ReCli.Incr(pkg.URLIdKey).Err()
	if err != nil {
		return -1, err
	}

	// encode global counter to base62
	id, err := pkg.ReCli.Get(pkg.URLIdKey).Int64()
	if err != nil {
		return -1, err
	}
	return id, nil
	// TODO should lock #1 end
}
