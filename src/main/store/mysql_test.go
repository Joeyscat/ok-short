package store

import (
	. "github.com/joeyscat/ok-short/common"
	"log"
	"testing"
)

const (
	driverName     = "mysql"
	dataSourceName = "root:123456@/ok-short?charset=utf8&parseTime=true"
)

func TestInsertLink(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	detail := &Link{
		OriginURL: "www",
		ShortCode: "s",
		CreatedBy: 0,
		CreatedAt: Now(),
		Exp:       30,
	}
	id, err := mySQL.InsertLink(detail)
	if err != nil {
		t.Error(err)
	}
	log.Println(id)
}

func TestQueryLink(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	var shortCode = "s"
	detail, err := mySQL.QueryLink(shortCode)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v\n", detail)
}

func TestMySQL_InsertLinkVisitedLog(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	visitedLog := &LinkVisitedLog{
		RemoteAddr: "test:test",
		ShortCode:  "test_sc",
		UA:         "test-user-agent",
		Cookie:     "test-cookie",
		VisitorId:  "0",
		VisitedAt:  Now(),
	}
	detail, err := mySQL.InsertLinkVisitedLog(visitedLog)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v\n", detail)
}

func TestMySQL_QueryLinkVisitedLog(t *testing.T) {
	mySQL := NewMySQL(driverName, dataSourceName)

	logs, err := mySQL.QueryLinkVisitedLog("K6")
	if err != nil {
		t.Error(err)
	}
	for _, visitedLog := range logs {
		log.Printf("%+v\n", visitedLog)
	}
}
