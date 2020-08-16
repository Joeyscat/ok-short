package admin

import (
	. "github.com/joeyscat/ok-short/common"
	. "github.com/joeyscat/ok-short/model"
	. "github.com/joeyscat/ok-short/store"
)

type LinkService struct {
}

func (l *LinkService) QueryLinkList(page, limit uint32) (*[]LinkRespData, uint32, error) {
	var list []Link
	var totalCount uint32
	offset := (page - 1) * limit
	MyDB.Offset(offset).Limit(limit).Find(&list).Offset(0).Count(&totalCount)

	var respList []LinkRespData
	for _, item := range list {
		l := fromLinkModel(&item)
		respList = append(respList, *l)
	}

	return &respList, totalCount, nil
}

func (l *LinkService) QueryLinkTraceList(page, limit uint32) (*[]LinkTraceRespData, uint32, error) {
	var list []LinkTrace
	var totalCount uint32
	offset := (page - 1) * limit
	MyDB.Offset(offset).Limit(limit).Find(&list).Offset(0).Count(&totalCount)

	var respList []LinkTraceRespData
	for _, item := range list {
		l := fromLinkTraceModel(&item)
		respList = append(respList, *l)
	}

	return &respList, totalCount, nil
}

func fromLinkModel(l *Link) *LinkRespData {
	return &LinkRespData{
		Sid:       l.Sid,
		URL:       LinkPrefix + l.Sc,
		Status:    l.Status,
		Name:      l.Name,
		OriginURL: l.OriginURL,
		CreatedAt: l.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func fromLinkTraceModel(l *LinkTrace) *LinkTraceRespData {
	return &LinkTraceRespData{
		Sid:       l.Sid,
		URL:       l.URL,
		UA:        l.UA,
		Ip:        l.Ip,
		CreatedAt: l.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
