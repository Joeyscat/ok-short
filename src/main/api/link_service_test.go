package api

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestShorten(t *testing.T) {
	service := LinkService{}

	url, err := service.Shorten("http://solo.oook.fun/articles/2019/06/29/1575110024441.html", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func TestLinkInfo(t *testing.T) {
	service := LinkService{}

	info, err := service.LinkInfo("XU")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}

func TestUnShorten(t *testing.T) {
	service := LinkService{}

	originUrl, err := service.UnShorten("XU")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(originUrl)
}

func TestAll(t *testing.T) {
	service := LinkService{}

	url, err := service.Shorten("http://solo.oook.fun/articles/2019/06/29/1575110024441.html", 30)
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(url, "/")
	sc := split[len(split)-1]
	fmt.Println(sc)

	info, err := service.LinkInfo(sc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)

	unShorten, err := service.UnShorten(sc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unShorten)
}
