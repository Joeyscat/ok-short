package main

import "net/http"

type Storage interface {
	Shorten(url string, exp uint32) (string, error)
	ShortURLInfo(eid string) (interface{}, error)
	UnShorten(eid string) (string, error)
	StoreVisitedLog(r *http.Request, shortCode string)
}
