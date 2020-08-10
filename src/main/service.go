package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Service struct {
	R RedisCli
}

// 将普通链接转为短链接
func (s *Service) Shorten(url string, exp int64) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := s.R.GenId()
	eid := Base62Encode(id)

	// store the url against this encoded id
	err = s.R.Cli.Set(fmt.Sprintf(ShortURLKey, eid), url,
		time.Minute*time.Duration(exp)).Err()

	if err != nil {
		return "", err
	}

	detail, err := json.Marshal(
		&URLDetail{
			URL:                 url,
			CreateAt:            time.Now().String(),
			ExpirationInMinutes: time.Duration(exp)})
	if err != nil {
		return "", err
	}

	// store the url detail against the encoded id
	err = s.R.Cli.Set(fmt.Sprintf(ShortURLDetailKey, eid), detail,
		time.Minute*time.Duration(exp)).Err()

	if err != nil {
		return "", err
	}

	return eid, nil
}

// ShortURLInfo returns the detail of the shorturl
func (s *Service) ShortURLInfo(eid string) (interface{}, error) {
	return s.R.UnShorten(ShortURLDetailKey, eid)
}

// UnShorten convert shorturl to url
func (s *Service) UnShorten(eid string) (string, error) {
	return s.R.UnShorten(ShortURLKey, eid)
}
