package admin

import (
	"errors"
	. "github.com/joeyscat/ok-short/common"
	. "github.com/joeyscat/ok-short/store"
	"log"
)

type LinkService struct {
	R RedisCli
	M MySQL
}

func (l *LinkService) QueryLinks(page, size uint32) (*[]Link, uint32, uint32, error) {
	links, count, total, err := l.M.GetLinkList(page*size, size)
	if err != nil {
		log.Println(err)
		return nil, 0, 0, StatusError{Code: 500, Err: errors.New("QueryLinks-查询失败")}
	}

	return links, count, total, nil
}
