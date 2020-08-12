package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
)

const (
	// URLIdKey redis自增主键所用的key
	URLIdKey = "next.url.id"
	// LinkKey 用于保存短链与原始链接的映射
	LinkKey = "link:%s:url"
)

type RedisCli struct {
	Cli *redis.Client
}

// NewRedisCli create a redis client
func NewRedisCli(addr string, password string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}

	return &RedisCli{Cli: c}
}

func (r *RedisCli) UnShorten(key, encodedId string) (string, error) {
	data, err := r.Cli.Get(fmt.Sprintf(key, encodedId)).Result()
	if err == redis.Nil {
		return "", StatusError{404, errors.New("Unknown Link")}
	} else if err != nil {
		return "", err
	} else {
		return data, nil
	}
}

func (r *RedisCli) GenId() (int64, error) {
	// TODO should lock #1 begin
	// increase the global counter
	err := r.Cli.Incr(URLIdKey).Err()
	if err != nil {
		return -1, err
	}

	// encode global counter to base62
	id, err := r.Cli.Get(URLIdKey).Int64()
	if err != nil {
		return -1, err
	}
	return id, nil
	// TODO should lock #1 end
}
