package api

import (
	"github.com/joeyscat/ok-short/common"
	"github.com/joeyscat/ok-short/model"
)

type Service interface {
	Shorten(url string, exp uint32) (string, error)
	LinkInfo(sc string) (*common.LinkRespData, error)
	UnShorten(sc string) (string, error)
	StoreVisitedLog(l *model.LinkTrace)
}
