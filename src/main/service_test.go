package main

import (
	"fmt"
	"log"
	"testing"
)

func TestShorten(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	service := Service{R: *redisCli}

	shortURL, err := service.Shorten("www.github.com", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shortURL)
}

func TestShortURLInfo(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	mySQL := NewMySQL(DriverName, DataSourceName)
	service := Service{R: *redisCli, M: *mySQL}

	info, err := service.ShortURLInfo("HK")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}

func TestUnShorten(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	service := Service{R: *redisCli}

	unShorten, err := service.UnShorten("z")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unShorten)
}

func TestAll(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	mySQL := NewMySQL(DriverName, DataSourceName)
	service := Service{R: *redisCli, M: *mySQL}

	shortURL, err := service.Shorten("www.github.com", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shortURL)

	info, err := service.ShortURLInfo(shortURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)

	unShorten, err := service.UnShorten(shortURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unShorten)
}
