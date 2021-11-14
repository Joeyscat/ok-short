package store

import (
	"context"

	"github.com/joeyscat/ok-short/internal/model"
)

type LinkStore interface {
	CreateLink(ctx context.Context, link *model.Link) error
	GetLinkBySc(ctx context.Context, sc string) (l *model.Link, err error)
	ListLinks(ctx context.Context, page, pageSize int64) (list []*model.Link, err error)
}
