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
	LinkCollectionName = "link"
)

type links struct {
	db *mongo.Database
}

func newLinks(ds *datastore) *links {
	return &links{db: ds.db}
}

func (s *links) CreateLink(ctx context.Context, link *model.Link) error {
	_, err := s.db.Collection(LinkCollectionName).InsertOne(ctx, link)

	return err
}

func (s *links) GetLinkBySc(ctx context.Context, sc string) (l *model.Link, err error) {
	l = &model.Link{}
	err = s.db.Collection(LinkCollectionName).FindOne(ctx, bson.M{"sc": sc}).Decode(l)

	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return
}

func (s *links) ListLinks(ctx context.Context, page, pageSize int64) (list []*model.Link, err error) {

	var opt *options.FindOptions
	if page > 0 && pageSize > 0 {
		skip := pageSize * (page - 1)
		opt = &options.FindOptions{Skip: &skip, Limit: &pageSize}
	}

	r, err := s.db.Collection(LinkCollectionName).Find(ctx, bson.M{}, opt)

	if err != nil {
		return nil, err
	}

	err = r.All(ctx, &list)
	return
}
