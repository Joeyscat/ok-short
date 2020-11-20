package model

import (
	"github.com/jinzhu/gorm"
	"github.com/joeyscat/ok-short/pkg/app"
)

const (
	StatueOpen  = "已启用"
	StatueClose = "已禁用"
)

// Link 短链结构体
type Link struct { // TODO
	*Model
	Sid       string // 业务标识
	Sc        string // 短链代码
	Status    string
	OriginURL string `gorm:"origin_url"` // 原始链接
	Exp       uint32
}

func (Link) TableName() string {
	return "ok_link"
}

// 追踪访问量等信息
type Trace struct {
}

type LinkSwagger struct {
	List  []*Link
	Pager *app.Pager
}

func (l Link) Create(db *gorm.DB) (*Link, error) {
	if err := db.Create(&l).Error; err != nil {
		return nil, err
	}
	return &l, nil
}

func (l Link) GetBySc(db *gorm.DB) (*Link, error) {
	var link Link
	db = db.Where("sc = ?", l.Sc)
	err := db.First(&link).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &link, nil
}

func (l Link) List(db *gorm.DB, pageOffset, pageSize int) ([]*Link, error) {
	var links []*Link
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Find(&links).Error; err != nil {
		return nil, err
	}

	return links, nil
}
