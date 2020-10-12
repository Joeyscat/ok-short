package dao

import (
	"github.com/joeyscat/ok-short/internel/model"
	"github.com/joeyscat/ok-short/pkg/app"
)

func (d *Dao) CreateLink(sc string, originalURL string, exp uint32) (*model.Link, error) {
	link := &model.Link{
		Sid:       app.Sid(),
		Sc:        sc,
		Status:    model.StatueOpen,
		Name:      "",
		OriginURL: originalURL,
		CreatedBy: 0,
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
