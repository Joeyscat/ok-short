package main

type Service interface {
	Shorten(url string, exp uint32) (string, error)
	LinkInfo(eid string) (interface{}, error)
	UnShorten(eid string) (string, error)
	StoreVisitedLog(l *LinkVisitedLog)
}
