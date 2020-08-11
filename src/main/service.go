package main

import (
	"fmt"
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

	//detail, err := json.Marshal(
	//	&URLDetail{
	//		URL:                 url,
	//		ShortCode:           eid,
	//		CreatedAt:            time.Now().String(),
	//		ExpirationInMinutes: time.Duration(exp)})
	//if err != nil {
	//	return "", err
	//}
	//
	//// store the url detail against the encoded id
	//err = s.R.Cli.Set(fmt.Sprintf(ShortURLDetailKey, eid), detail,
	//	expiration).Err()
	//
	//if err != nil {
	//	return "", err
	//}
	detail := &URLDetail{
		URL:                 url,
		ShortCode:           eid,
		CreatedBy:           0,
		CreatedAt:           time.Now(),
		ExpirationInMinutes: exp}
	_, err = s.M.InsertURLDetail(detail)
	if err != nil {
		return "", err
	}

	return eid, nil
}

// ShortURLInfo returns the detail of the shortURL
func (s *Service) ShortURLInfo(eid string) (interface{}, error) {
	return s.M.QueryUrlDetail(eid)
}

// UnShorten convert shortURL to url
func (s *Service) UnShorten(eid string) (string, error) {
	return s.R.UnShorten(ShortURLKey, eid)
}
