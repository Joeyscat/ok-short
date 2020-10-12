package dao

import (
	"github.com/joeyscat/ok-short/internel/model"
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
