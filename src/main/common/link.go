package common

import "time"

// Link 短链结构体
type Link struct { // TODO
	Id        uint32    `json:"-"`            // 数据库自增ID，不得传递到前端
	Sid       string    `gorm:"-" json:"sid"` // 业务标识
	URL       string    `gorm:"-" json:"url"` // 短链
	Enabled   bool      `gorm:"-" json:"enabled"`
	Group     Group     `gorm:"-" json:"group"` // 分组
	Name      string    `gorm:"-" json:"name"`
	OriginURL string    `json:"origin_url"` // 原始链接
	Trace     Trace     `gorm:"-" json:"trace"`
	ShortCode string    `json:"short_code"`
	CreatedBy uint32    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Exp       uint32    `gorm:"column:expiration_in_minutes" json:"-"`
}

func (Link) TableName() string {
	return "ok_link"
}

// 短链分组
type Group struct {
	Id   uint32 `json:"-"`
	Sid  string `json:"sid"`
	Name string `json:"name"`
}

// 追踪访问量等信息
type Trace struct {
}

type LinkVisitedLog struct {
	Id         string    `json:"id"`
	RemoteAddr string    `json:"remote_addr"`
	ShortCode  string    `json:"short_code"`
	UA         string    `json:"ua"`
	Cookie     string    `json:"cookie"`
	VisitorId  string    `json:"visitor_id"`
	VisitedAt  time.Time `json:"visited_at"`
}

func (LinkVisitedLog) TableName() string {
	return "ok_link_visited_log"
}
