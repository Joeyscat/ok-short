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

	detail := &GenURLDetail{
		URL:                 "www",
		ShortCode:           "s",
		CreatedBy:           0,
		CreatedAt:           time.Now(),
		ExpirationInMinutes: 30,
	}
	id, err := mySQL.InsertGenURLDetail(detail)
	if err != nil {
		t.Error(err)
	}
	log.Println(id)
}

func TestQueryURLDetail(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	var shortCode = "s"
	detail, err := mySQL.QueryGenURLDetail(shortCode)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v\n", detail)
}

func TestMySQL_InsertShortURLVisitedLog(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	visitedLog := &ShortURLVisitedLog{
		RemoteAddr: "test:test",
		ShortCode:  "test_sc",
		UA:         "test-user-agent",
		Cookie:     "test-cookie",
		VisitorId:  "0",
		VisitedAt:  time.Now(),
	}
	detail, err := mySQL.InsertShortURLVisitedLog(visitedLog)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v\n", detail)
}

func TestMySQL_QueryShortURLVisitedLog(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	logs, err := mySQL.QueryShortURLVisitedLog("K6")
	if err != nil {
		t.Error(err)
	}
	for _, visitedLog := range logs {
		log.Printf("%+v\n", visitedLog)
	}
}
