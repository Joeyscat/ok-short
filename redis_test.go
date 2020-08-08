package main

import (
	"fmt"
	"log"
	"testing"
)

func ShortenTest(t *testing.T) {
	redisCli := NewRedisCli("localhost:6379", "", 0)

	shortURL, err := redisCli.Shorten("www.github.com", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shortURL)
}
