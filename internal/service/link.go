package service

import (
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
)

const (
	// URLIdKey redis自增主键所用的key
	URLIdKey = "next.url.id"
)

type CreateLinkRequest struct {
	URL                 string `json:"url" form:"url" validate:"required,min=20,max=200"`
	ExpirationInMinutes uint32 `json:"expiration_in_minutes" form:"created_by" validate:"min=0,max=1440"`
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

	//expiration := time.Minute * time.Duration(param.ExpirationInMinutes)
	// TODO 定时清理过期数据
	_, err = model.CreateLink(sc, param.URL, param.ExpirationInMinutes)

	return url, err
}

func (svc *Service) UnShorten(param *RedirectLinkRequest) (string, error) {
	//return global.Redis.Get(fmt.Sprintf(LinkKey, param.Sc)).Result()
	return model.GetLinkBySc(param.Sc)
}

func (svc *Service) GetLink(param *GetLinkRequest) (*model.Link, error) {
	return model.GetLinkDetailBySc(param.Sc)
}

func (svc *Service) GetLinkList(pager *app.Pager) ([]*model.Link, error) {
	return model.GetLinkList(int64(pager.Page), int64(pager.PageSize))
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
