package ok_short

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joeyscat/ok-short/internel/pkg"
	. "github.com/joeyscat/ok-short/internel/pkg/common"
	"github.com/joeyscat/ok-short/internel/pkg/model"
	"log"
	"time"
)

type LinkService struct {
}

// 将普通链接转为短链接
func (s *LinkService) Shorten(originURL string, exp uint32) (string, error) {
	// 生成ID，并进行62进制编码
	id, err := genId()
	sc := Base62Encode(id)
	url := pkg.LinkPrefix + sc

	expiration := time.Minute * time.Duration(exp)
	err = pkg.ReCli.Set(fmt.Sprintf(pkg.LinkKey, sc), originURL, expiration).Err()
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
	pkg.MyDB.Create(detail)

	if pkg.MyDB.NewRecord(detail) {
		return "", StatusError{Code: LinkCreateFail, Err: errors.New(BSText(LinkCreateFail))}
	}

	return url, nil
}

// LinkInfo returns the detail of the link
func (s *LinkService) LinkInfo(sc string) (*LinkRespData, error) {
	var link model.Link
	pkg.MyDB.Where("sc = ?", sc).Last(&link)

	if link.OriginURL == "" {
		return nil, StatusError{
			Code: LinkNotExists,
			Err:  errors.New(BSText(LinkNotExists)),
		}
	}

	return &LinkRespData{
		Sid:       link.Sid,
		URL:       pkg.LinkPrefix + link.Sc,
		Status:    link.Status,
		Name:      link.Name,
		OriginURL: link.OriginURL,
		CreatedAt: link.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// UnShorten 将短链还原为原始链接
func (s *LinkService) UnShorten(sc string) (string, error) {
	data, err := pkg.ReCli.Get(fmt.Sprintf(pkg.LinkKey, sc)).Result()

	if err == redis.Nil {
		return "", StatusError{Code: LinkNotExists, Err: errors.New(BSText(LinkNotExists))}
	} else if err != nil {
		return "", err
	} else {
		return data, nil
	}
}

func (s *LinkService) StoreVisitedLog(l *model.LinkTrace) {

	pkg.MyDB.Create(l)

	if pkg.MyDB.NewRecord(l) {
		log.Println("插入短链访问记录失败")
	}
}

func genId() (int64, error) {
	// TODO should lock #1 begin
	// increase the global counter
	err := pkg.ReCli.Incr(pkg.URLIdKey).Err()
	if err != nil {
		return -1, err
	}

	// encode global counter to base62
	id, err := pkg.ReCli.Get(pkg.URLIdKey).Int64()
	if err != nil {
		return -1, err
	}
	return id, nil
	// TODO should lock #1 end
}
