package main

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	// URLIDKEY is global counter
	URLIDKEY = "next.url.id"
	// ShortlinkKey mapping the shortlink to the url
	ShortlinkKey = "shortlink:%s:url"
	// URLHashKey mapping the hash of the url to the shortlink
	URLHashKey = "urlhash:%s:url"
	// ShortlinkDetailKey mapping the shortlink to the detail of url
	ShortlinkDetailKey = "shortlink:%s:detail"
)

// RedisCli contains a redis client
type RedisCli struct {
	Cli *redis.Client
}

// URLDetail contains the detail of the shortlink
type URLDetail struct {
	URL                 string        `json:"url"`
	CreateAt            string        `json:"created_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

// NewRedisCli create a redis client
func NewRedisCli(addr string, passwd string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})

	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}

	return &RedisCli{Cli: c}
}

// Shorten convert url to shortlink
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	// convert url to sha1 hash
	h := toSha1(url)

	// fetch it if the url is cached
	data, err := r.Cli.Get(fmt.Sprintf(URLHashKey, h)).Result()
	if err == redis.Nil {
		// not exist, nothing to do
	} else if err != nil {
		return "", err
	} else {
		if data == "{}" {
			// expiration, nothing to do
		} else {
			return data, nil
		}
	}
	// TODO should lock #1 begin
	// increase the global counter
	err = r.Cli.Incr(URLIDKEY).Err()
	if err != nil {
		return "", err
	}

	// encode global counter to base62
	id, err := r.Cli.Get(URLIDKEY).Int64()
	eid := Base62Encode(id)
	// TODO should lock #1 end

	// store the url against this encoded id
	err = r.Cli.Set(fmt.Sprintf(ShortlinkKey, eid), url,
		time.Minute*time.Duration(exp)).Err()

	if err != nil {
		return "", err
	}

	// store the url against the hash of it
	err = r.Cli.Set(fmt.Sprintf(URLHashKey, h), eid,
		time.Minute*time.Duration(exp)).Err()

	if err != nil {
		return "", err
	}

	detail, err := json.Marshal(
		&URLDetail{
			URL:                 url,
			CreateAt:            time.Now().String(),
			ExpirationInMinutes: time.Duration(exp)})
	if err != nil {
		return "", err
	}

	// store the url detail aginst the encoded id
	err = r.Cli.Set(fmt.Sprintf(ShortlinkDetailKey, eid), detail,
		time.Minute*time.Duration(exp)).Err()

	if err != nil {
		return "", err
	}

	return eid, nil
}

// ShortlinkInfo returns the detail of the shortlink
func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, error) {
	data, err := r.Cli.Get(fmt.Sprintf(ShortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", StatusError{404, errors.New("Unknown short URL")}
	} else if err != nil {
		return "", err
	} else {
		return data, nil
	}
}

// UnShorten convert shortlink to url
func (r *RedisCli) UnShorten(eid string) (string, error) {
	data, err := r.Cli.Get(fmt.Sprintf(ShortlinkKey, eid)).Result()
	if err == redis.Nil {
		return "", StatusError{404, errors.New("Unknown short URL")}
	} else if err != nil {
		return "", err
	} else {
		return data, nil
	}
}

func toSha1(str string) string {
	var (
		sha = sha1.New()
	)
	return string(sha.Sum([]byte(str)))
}
