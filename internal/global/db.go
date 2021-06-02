package global

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joeyscat/ok-short/pkg/setting"
	"github.com/qiniu/qmgo"
	"strings"
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
	instances := strings.Join(s.Addr, ",")
	auth := ""
	if s.User != "" && s.Password != "" {
		auth = fmt.Sprintf("%s:%s@", s.User, s.Password)
	}
	uri := fmt.Sprintf("mongodb://%s%s/%s", auth, instances, s.AuthDB)
	fmt.Println(uri)

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: uri})

	return client, err
}
