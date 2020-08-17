package app

import (
	"github.com/joeyscat/ok-short/internel/app/ok-short"
	"github.com/joeyscat/ok-short/internel/app/ok-short-admin"
	"github.com/joeyscat/ok-short/internel/pkg"
)

type Context struct {
	API   ok_short.Service
	ADMIN ok_short_admin.Service
	Port  int
}

func GetContext() *Context {
	return &Context{
		API: &ok_short.LinkService{},
		ADMIN: ok_short_admin.Service{
			AuthorService: ok_short_admin.AuthorService{},
			LinkService:   ok_short_admin.LinkService{},
		},
		Port: pkg.ServerPort}
}
