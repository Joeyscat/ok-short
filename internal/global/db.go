package global

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/joeyscat/ok-short/pkg/setting"
	"github.com/qiniu/qmgo"
)

var (
	MongoOkShortDB      *qmgo.Database
	MongoLinksColl      *qmgo.Collection
	MongoLinksTraceColl *qmgo.Collection
	MongoAuthsColl      *qmgo.Collection

	Redis *redis.Client
)

func NewRedis(s *setting.RedisSettingS) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     s.Addr,
		Password: s.Password,
		DB:       s.DB,
	})
	return client
}

func NewMongoDB(s *setting.MongoDBSettingS) (*qmgo.Client, error) {
	uri := s.URI
	fmt.Println(uri)

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: uri})

	return client, err
}
