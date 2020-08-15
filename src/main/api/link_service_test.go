package api

import (
	"fmt"
	"log"
	"testing"
)

func TestShorten(t *testing.T) {
	service := LinkService{}

	url, err := service.Shorten("www.github.com", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func TestLinkInfo(t *testing.T) {
	service := LinkService{}

	info, err := service.LinkInfo("Gq")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}

func TestUnShorten(t *testing.T) {
	service := LinkService{}

	originUrl, err := service.UnShorten("z")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(originUrl)
}

func TestAll(t *testing.T) {
	service := LinkService{}

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
