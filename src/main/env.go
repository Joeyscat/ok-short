package main

import (
	"github.com/joeyscat/ok-short/admin"
	"github.com/joeyscat/ok-short/api"
	. "github.com/joeyscat/ok-short/store"
)

type Env struct {
	API   api.Service
	ADMIN admin.Service
	port  int
}

func getEnv() *Env {
	return &Env{
		API: &api.LinkService{},
		ADMIN: admin.Service{
			AuthorService:  admin.AuthorService{},
			LinkService:    admin.LinkService{},
			VisitorService: admin.VisitorService{},
		},
		port: ServerPort}
}
