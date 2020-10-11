package model

import (
	"github.com/jinzhu/gorm"
	"github.com/joeyscat/ok-short/pkg/app"
)

// Link 短链结构体
type Link struct { // TODO
	*Model
	Sid       string // 业务标识
	Sc        string // 短链代码
	Status    string
	Group     Group `gorm:"-"` // 分组
	Name      string
	OriginURL string `gorm:"origin_url"` // 原始链接
	PV        Trace  `gorm:"-"`
	CreatedBy uint32
	Exp       uint32
}

func (Link) TableName() string {
	return "ok_link"
}

// 短链分组
type Group struct {
	Id   uint32 `gorm:"-"`
	Sid  string `gorm:"sid"`
	Name string `gorm:"name"`
}

// 追踪访问量等信息
type Trace struct {
}

type LinkTrace struct {
	gorm.Model
	Sid    string
	URL    string `gorm:"url"`
	Ip     string `gorm:"ip"`
	UA     string `gorm:"ua"`
	Cookie string `gorm:"cookie"`
}

func (LinkTrace) TableName() string {
	return "ok_link_trace"
}

type LinkSwagger struct {
	List  []*Link
	Pager *app.Pager
}

func (l *Link) Create(db *gorm.DB) error {
	return db.Create(&l).Error
}

func (l *Link) Find(db *gorm.DB) *Link {
	db.Find(&l)
	return l
}
