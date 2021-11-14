package store

import (
	"context"

	"github.com/joeyscat/ok-short/internal/model"
)

type LinkTraceStore interface {
	CreateLinkTrace(ctx context.Context, lt *model.LinkTrace) error
	ListLinkTrace(ctx context.Context, page, pageSize int64) (lt []*model.LinkTrace, err error)
	ListLinkTraceBySc(ctx context.Context, sc string, page, pageSize int64) (lt []*model.LinkTrace, err error)
}
