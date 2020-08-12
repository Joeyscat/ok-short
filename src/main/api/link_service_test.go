package api

import (
	"fmt"
	"github.com/joeyscat/ok-short"
	. "github.com/joeyscat/ok-short/store"
	"log"
	"testing"
)

func TestShorten(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	service := LinkService{R: *redisCli}

	url, err := service.Shorten("www.github.com", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func TestLinkInfo(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	mySQL := NewMySQL(main.DriverName, main.DataSourceName)
	service := LinkService{R: *redisCli, M: *mySQL}

	info, err := service.LinkInfo("HK")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}

func TestUnShorten(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	service := LinkService{R: *redisCli}

	originUrl, err := service.UnShorten("z")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(originUrl)
}

func TestAll(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	mySQL := NewMySQL(main.DriverName, main.DataSourceName)
	service := LinkService{R: *redisCli, M: *mySQL}

	eid, err := service.Shorten("www.github.com", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(eid)

	info, err := service.LinkInfo(eid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)

	unShorten, err := service.UnShorten(eid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unShorten)
}
