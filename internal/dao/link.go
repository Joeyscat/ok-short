package dao

import (
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
)

func (d *Dao) CreateLink(sc string, originalURL string, exp uint32) (*model.Link, error) {
	link := &model.Link{
		Sid:       app.Sid(),
		Sc:        sc,
		OriginURL: originalURL,
		Exp:       exp,
	}
	return link.Create(d.engine)
}

func (d *Dao) GetLink(sc string) (*model.Link, error) {
	link := &model.Link{
		Sc: sc,
	}
	return link.GetBySc(d.engine)
}

func (d *Dao) GetLinkList(page, pageSize int) ([]*model.Link, error) {
	link := &model.Link{}
	return link.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}
