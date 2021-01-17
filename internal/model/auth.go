package model

import (
	"context"
	"github.com/joeyscat/ok-short/global"
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
)

type Auth struct {
	field.DefaultField `bson:",inline"`

	AppKey    string `bson:"app_key"`
	AppSecret string `bson:"app_secret"`
}

func CreateAuth(appKey, appSecret string) (*Auth, error) {
	a := &Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	_, err := global.MongoAuthsColl.InsertOne(context.Background(), a)
	return a, err
}

func GetAuth(appKey, appSecret string) (a *Auth, err error) {
	a = &Auth{}
	// 覆盖索引
	err = global.MongoAuthsColl.Find(context.Background(),
		bson.M{"app_key": appKey, "app_secret": appSecret}).
		Select(bson.M{"_id": 0, "app_key": 1}).One(a)
	return a, err
}
