package dao

import (
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
)

func (d *Dao) CreateLinkTrace(sc, url, ip, ua, cookies string) (*model.LinkTrace, error) {
	l := &model.LinkTrace{
		Model:  nil,
		Sc:     sc,
		URL:    url,
		Ip:     ip,
		UA:     ua,
		Cookie: cookies,
	}
	return l.Create(d.engine)
}

func (d *Dao) GetLinkTrace(sc string, page, pageSize int) ([]*model.LinkTrace, error) {
	trace := &model.LinkTrace{
		Sc: sc,
	}
	return trace.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}
