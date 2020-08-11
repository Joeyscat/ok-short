package main

import (
	"log"
	"testing"
	"time"
)

const (
	driverName     = "mysql"
	dataSourceName = "root:123456@/ok-short?charset=utf8&parseTime=true"
)

func TestInsertURLDetail(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	detail := URLDetail{
		URL:                 "www",
		ShortCode:           "s",
		CreatedBy:           0,
		CreatedAt:           time.Now(),
		ExpirationInMinutes: 30,
	}
	id, err := mySQL.InsertURLDetail(&detail)
	if err != nil {
		t.Error(err)
	}
	log.Println(id)
}

func TestQueryURLDetail(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	var shortCode = "s"
	detail, err := mySQL.QueryUrlDetail(shortCode)
	if err != nil {
		t.Error(err)
	}
	log.Println(detail)
}
