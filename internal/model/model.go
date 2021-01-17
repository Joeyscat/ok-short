package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func NewRedis(s *setting.RedisSettingS) *redis.Client {
//	client := redis.NewClient(&redis.Options{
//		Addr:     s.Addr,
//		Password: s.Password,
//		DB:       s.DB,
//	})
//	return client
//}
//
//func NewMongoDB(s *setting.MongoDBSettingS) (*qmgo.Client, error) {
//	instances := strings.Join(s.Addr, ",")
//	auth := ""
//	if s.User != "" && s.Password != "" {
//		auth = fmt.Sprintf("%s:%s@", s.User, s.Password)
//	}
//	uri := fmt.Sprintf("mongodb://%s%s/%s", auth, instances, s.AuthDB)
//	fmt.Println(uri)
//
//	ctx := context.Background()
//	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: uri})
//
//	return client, err
//}
