package main

import (
	"fmt"
	"log"
	"testing"
)

func TestShorten(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	service := Service{R: *redisCli}

	shortURL, err := service.Shorten("www.github.com", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shortURL)
}

func TestShortURLInfo(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)
	service := Service{R: *redisCli}
	info, err := service.ShortURLInfo("z")
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
