package admin

import (
	. "github.com/joeyscat/ok-short/common"
	. "github.com/joeyscat/ok-short/store"
)

type LinkService struct {
}

func (l *LinkService) QueryLinks(page, size uint32) (*[]Link, uint32, uint32, error) {
	var links []Link
	var count uint32
	var total uint32
	offset := (page - 1) * size
	MyDB.Offset(&offset).Limit(&size).Find(&links).Count(&count)

	return &links, count, total, nil
}
