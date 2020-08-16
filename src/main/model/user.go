package model

import (
	"github.com/jinzhu/gorm"
)

// 管理员
type User struct {
	gorm.Model
	Name      string
	Password  string
	Email     string `gorm:"type:varchar(100);unique_index"`
	AvatarURL string
}

// 将 User 的表名设置为 `ok_link_admin_user`
func (User) TableName() string {
	return "ok_link_admin_user"
}
