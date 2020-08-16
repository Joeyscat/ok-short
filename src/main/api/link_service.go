package api

import (
	"errors"
	"fmt"
	. "github.com/joeyscat/ok-short/common"
	"github.com/joeyscat/ok-short/model"
	. "github.com/joeyscat/ok-short/store"
	"log"
	"time"
)

type LinkService struct {
}

// 将普通链接转为短链接
func (s *LinkService) Shorten(originURL string, exp uint32) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := ReCli.GenId()
	sc := Base62Encode(id)
	url := LinkPrefix + sc

	expiration := time.Minute * time.Duration(exp)
	err = ReCli.Cli.Set(fmt.Sprintf(LinkKey, sc), originURL, expiration).Err()
	if err != nil {
		return "", err
	}
	// 存储原链接与短链接代码的映射
	detail := &model.Link{
		Sid:       Sid(),
		Sc:        sc,
		Status:    "已启用",
		Name:      "",
		OriginURL: originURL,
		CreatedBy: 0,
		Exp:       exp,
	}
	// TODO 定时清理过期数据
	MyDB.Create(detail)

	if MyDB.NewRecord(detail) {
		return "", StatusError{Code: LinkCreateFail, Err: errors.New(BSText(LinkCreateFail))}
	}

	return url, nil
}

// LinkInfo returns the detail of the link
func (s *LinkService) LinkInfo(sc string) (*LinkRespData, error) {
	var link model.Link
	MyDB.Where("sc = ?", sc).Last(&link)

	if link.OriginURL == "" {
		return nil, StatusError{
			Code: LinkNotExists,
			Err:  errors.New(BSText(LinkNotExists)),
		}
	}

	return &LinkRespData{
		Sid:       link.Sid,
		URL:       LinkPrefix + link.Sc,
		Status:    link.Status,
		Name:      link.Name,
		OriginURL: link.OriginURL,
		CreatedAt: link.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// UnShorten 将短链还原为原始链接
func (s *LinkService) UnShorten(sc string) (string, error) {
	return ReCli.UnShorten(LinkKey, sc)
}

func (s *LinkService) StoreVisitedLog(l *model.LinkTrace) {

	MyDB.Create(l)

	if MyDB.NewRecord(l) {
		log.Println("插入短链访问记录失败")
	}
}
