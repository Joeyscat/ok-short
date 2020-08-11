package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Service struct {
	R RedisCli
	M MySQL
}

// 将普通链接转为短链接
func (s *Service) Shorten(url string, exp uint32) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := s.R.GenId()
	eid := Base62Encode(id)

	expiration := time.Minute * time.Duration(exp)
	// store the url against this encoded id
	err = s.R.Cli.Set(fmt.Sprintf(ShortURLKey, eid), url,
		expiration).Err()

	if err != nil {
		return "", err
	}

	detail := &GenURLDetail{
		URL:                 url,
		ShortCode:           eid,
		CreatedBy:           0,
		CreatedAt:           time.Now(),
		ExpirationInMinutes: exp}
	// TODO 定时清理过期数据
	_, err = s.M.InsertGenURLDetail(detail)
	if err != nil {
		return "", err
	}

	return eid, nil
}

// ShortURLInfo returns the detail of the shortURL
func (s *Service) ShortURLInfo(eid string) (interface{}, error) {
	return s.M.QueryGenURLDetail(eid)
}

// UnShorten convert shortURL to url
func (s *Service) UnShorten(eid string) (string, error) {
	return s.R.UnShorten(ShortURLKey, eid)
}

func (s *Service) StoreVisitedLog(r *http.Request, shortCode string) {
	visitedLog := GetVisitedLog(r, shortCode)

	_, err := s.M.InsertShortURLVisitedLog(visitedLog)
	if err != nil {
		log.Println(err)
	}
}
