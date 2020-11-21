package dao

import (
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
)

func (d *Dao) CreateLinkTrace(l *model.LinkTrace) (*model.LinkTrace, error) {
	return l.Create(d.engine)
}

func (d *Dao) GetLinkTrace(sc string, page, pageSize int) ([]*model.LinkTrace, error) {
	trace := &model.LinkTrace{
		Sc: sc,
	}
	return trace.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}
