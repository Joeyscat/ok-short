package mongo

import (
	"context"
	"errors"

	"github.com/joeyscat/ok-short/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	LinkTraceCollectionName = "link_trace"
)

type linkTraces struct {
	db *mongo.Database
}

func newLinkTraces(ds *datastore) *linkTraces {
	return &linkTraces{db: ds.db}
}

func (s *linkTraces) CreateLinkTrace(ctx context.Context, lt *model.LinkTrace) error {
	_, err := s.db.Collection(LinkTraceCollectionName).InsertOne(ctx, lt)

	return err
}

func (s *linkTraces) ListLinkTrace(ctx context.Context, page, pageSize int64) (lt []*model.LinkTrace, err error) {
	var opt *options.FindOptions
	if page > 0 && pageSize > 0 {
		skip := pageSize * (page - 1)
		opt = &options.FindOptions{Skip: &skip, Limit: &pageSize}
	}

	r, err := s.db.Collection(LinkTraceCollectionName).Find(ctx, bson.M{}, opt)

	if err != nil {
		return nil, err
	}

	err = r.All(ctx, &lt)
	return
}

func (s *linkTraces) ListLinkTraceBySc(ctx context.Context, sc string, page, pageSize int64) (lt []*model.LinkTrace, err error) {
	if sc == "" {
		return nil, errors.New("sc could not be empty")
	}
	var opt *options.FindOptions
	if page > 0 && pageSize > 0 {
		skip := pageSize * (page - 1)
		opt = &options.FindOptions{Skip: &skip, Limit: &pageSize}
	}

	r, err := s.db.Collection(LinkTraceCollectionName).Find(ctx, bson.M{"sc": sc}, opt)

	if err != nil {
		return nil, err
	}

	err = r.All(ctx, &lt)
	return
}
