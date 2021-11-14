package v1

import (
	"context"
	"fmt"

	"github.com/joeyscat/ok-short/internal/global"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/store"
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

type LinkSrv interface {
	CreateLink(ctx context.Context, param *CreateLinkRequest) (string, error)
	UnShorten(ctx context.Context, param *RedirectLinkRequest) (string, error)
	GetLink(ctx context.Context, param *GetLinkRequest) (*model.Link, error)
	ListLinks(ctx context.Context, pager *app.Pager) ([]*model.Link, error)
}

type linkService struct {
	store store.Factory
}

var _ LinkSrv = (*linkService)(nil)

func newLinks(srv *service) *linkService {
	return &linkService{store: srv.store}
}

func (s *linkService) CreateLink(ctx context.Context, param *CreateLinkRequest) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := genId()
	if err != nil {
		return "", err
	}
	sc := app.Base62Encode(id)
	url := global.AppSetting.LinkPrefix + sc

	//expiration := time.Minute * time.Duration(param.ExpirationInMinutes)
	// TODO 定时清理过期数据
	l := &model.Link{
		Sc:        sc,
		OriginURL: param.URL,
		Exp:       param.ExpirationInMinutes,
	}
	err = s.store.Links().CreateLink(ctx, l)

	return url, err
}
func (s *linkService) UnShorten(ctx context.Context, param *RedirectLinkRequest) (string, error) {
	//return global.Redis.Get(fmt.Sprintf(LinkKey, param.Sc)).Result()
	l, err := s.store.Links().GetLinkBySc(ctx, param.Sc)
	if err != nil {
		return "", err
	}
	if l == nil {
		return "", nil
	}

	return l.OriginURL, nil
}

func (s *linkService) GetLink(ctx context.Context, param *GetLinkRequest) (*model.Link, error) {
	l, err := s.store.Links().GetLinkBySc(ctx, param.Sc)
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, fmt.Errorf("short code [%s] not exist", param.Sc)
	}
	return l, err
}

func (s *linkService) ListLinks(ctx context.Context, pager *app.Pager) ([]*model.Link, error) {
	return s.store.Links().ListLinks(ctx, int64(pager.Page), int64(pager.PageSize))
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
