package model

import "github.com/jinzhu/gorm"

type LinkTrace struct {
	*Model
	Sc     string
	URL    string `gorm:"url"`
	Ip     string `gorm:"ip"`
	UA     string `gorm:"ua"`
	Cookie string `gorm:"cookie"`
}

func (LinkTrace) TableName() string {
	return "ok_link_trace"
}

func (lt LinkTrace) Create(db *gorm.DB) (*LinkTrace, error) {
	if err := db.Create(&lt).Error; err != nil {
		return nil, err
	}
	return &lt, nil
}

func (lt LinkTrace) Get(db *gorm.DB) (*LinkTrace, error) {
	if err := db.Find(&lt).Error; err != nil {
		return nil, err
	}
	return &lt, nil
}

func (lt LinkTrace) List(db *gorm.DB, pageOffset, pageSize int) ([]*LinkTrace, error) {
	var traces []*LinkTrace
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if lt.Sc != "" {
		db = db.Where("sc = ?", lt.Sc)
	}
	if lt.URL != "" {
		db = db.Where("url = ?", lt.URL)
	}
	if err = db.Find(&traces).Error; err != nil {
		return nil, err
	}

	return traces, nil
}

func (lt LinkTrace) ListByURL(db *gorm.DB, url string) ([]*LinkTrace, error) {
	return nil, nil
}
