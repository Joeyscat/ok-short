package v1

import (
	"context"

	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/store"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/labstack/echo/v4"
)

type GetLinkTraceRequest struct {
	Sc string `json:"sc" form:"sc" binding:"required"`
}

type LinkTraceSrv interface {
	CreateLinkTrace(ctx context.Context, sc, url string, c echo.Context) error
	// GetLinkTrace 获取某短链访问记录，返回一个LinkTrace数组
	GetLinkTrace(ctx context.Context, param *GetLinkTraceRequest, pager *app.Pager) ([]*model.LinkTrace, error)
	// ListLinkTrace 获取多个短链访问记录
	ListLinkTrace(ctx context.Context, pager *app.Pager) ([]*model.LinkTrace, error)
}

type linkTraceService struct {
	store store.Factory
}

var _ LinkTraceSrv = (*linkTraceService)(nil)

func newLinkTraces(srv *service) *linkTraceService {
	return &linkTraceService{store: srv.store}
}

func (s *linkTraceService) CreateLinkTrace(ctx context.Context, sc, url string, c echo.Context) error {
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

	return s.store.LinkTraces().CreateLinkTrace(ctx, lt)
}

func (s *linkTraceService) GetLinkTrace(ctx context.Context, param *GetLinkTraceRequest, pager *app.Pager) ([]*model.LinkTrace, error) {
	return s.store.LinkTraces().ListLinkTraceBySc(ctx, param.Sc, int64(pager.Page), int64(pager.PageSize))
}

func (s *linkTraceService) ListLinkTrace(ctx context.Context, pager *app.Pager) ([]*model.LinkTrace, error) {
	return s.store.LinkTraces().ListLinkTrace(ctx, int64(pager.Page), int64(pager.PageSize))
}
