package main

import (
	"fmt"
	"log"
	"testing"
)

func TestShorten(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)

	shortURL, err := redisCli.Shorten("www.github.com", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shortURL)
}

func TestShortlinkInfo(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)

	info, err := redisCli.ShortlinkInfo("7")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}

func TestUnShorten(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)

	unShorten, err := redisCli.UnShorten("7")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unShorten)
}
