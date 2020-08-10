package main

type Storage interface {
	Shorten(url string, exp int64) (string, error)
	ShortURLInfo(eid string) (interface{}, error)
	UnShorten(eid string) (string, error)
}
