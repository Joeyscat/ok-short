package v1

import (
	"github.com/joeyscat/ok-short/internal/store"
)

type Service interface {
	Links() LinkSrv
	LinkTraces() LinkTraceSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Links() LinkSrv {
	return newLinks(s)
}

func (s *service) LinkTraces() LinkTraceSrv {
	return newLinkTraces(s)
}
