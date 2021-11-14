package global

import (
	"github.com/go-redis/redis"
	"github.com/joeyscat/ok-short/pkg/setting"
)

var (
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
