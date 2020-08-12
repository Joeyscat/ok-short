package main

import (
	"fmt"
	"log"
	"time"
)

type LinkService struct {
	R RedisCli
	M MySQL
}

// Link 短链结构体
type Link struct { // TODO
	Id        uint32    // 数据库自增ID，不得传递到前端
	Sid       string    `json:"sid"` // 业务标识
	URL       string    `json:"url"` // 短链
	Enabled   bool      `json:"enabled"`
	Group     Group     `json:"group"` // 分组
	Name      string    `json:"name"`
	OriginURL string    `json:"origin_url"` // 原始链接
	Trace     Trace     `json:"stats"`
	ShortCode string    `json:"short_code"`
	CreatedBy uint32    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	Exp       uint32    `json:"exp"`
}

// 短链分组
type Group struct {
	Id   uint32
	Sid  string `json:"sid"`
	Name string `json:"name"`
}

// 追踪访问量等信息
type Trace struct {
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
