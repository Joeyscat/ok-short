package api

import (
	"fmt"
	. "github.com/joeyscat/ok-short/common"
	. "github.com/joeyscat/ok-short/store"
	"log"
	"time"
)

type LinkService struct {
	R RedisCli
	M MySQL
}

// 将普通链接转为短链接
func (s *LinkService) Shorten(originURL string, exp uint32) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := s.R.GenId()
	eid := Base62Encode(id)

	expiration := time.Minute * time.Duration(exp)
	// 存储原链接与短链接代码的映射
	err = s.R.Cli.Set(fmt.Sprintf(LinkKey, eid), originURL, expiration).Err()
	if err != nil {
		return "", err
	}

	detail := &Link{
		OriginURL: originURL,
		ShortCode: eid,
		CreatedBy: 0,
		CreatedAt: Now(),
		Exp:       exp,
	}
	// TODO 定时清理过期数据
	_, err = s.M.InsertLink(detail)
	if err != nil {
		return "", err
	}

	return eid, nil
}

// LinkInfo returns the detail of the link
func (s *LinkService) LinkInfo(eid string) (interface{}, error) {
	return s.M.QueryLink(eid)
}

// UnShorten 将短链还原为原始链接
func (s *LinkService) UnShorten(eid string) (string, error) {
	return s.R.UnShorten(LinkKey, eid)
}

func (s *LinkService) StoreVisitedLog(l *LinkVisitedLog) {
	_, err := s.M.InsertLinkVisitedLog(l)
	if err != nil {
		log.Println(err)
	}
}
