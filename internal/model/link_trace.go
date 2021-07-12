package model

import (
	"context"
	"github.com/joeyscat/ok-short/internal/global"
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkTrace struct {
	field.DefaultField `bson:",inline"`

	Sc     string `bson:"sc"`
	URL    string `bson:"url"`
	Ip     string `bson:"ip"`
	UA     string `bson:"ua"`
	Cookie string `bson:"cookie"`
}

func CreateLinkTrace(lt *LinkTrace) error {
	_, err := global.MongoLinksTraceColl.InsertOne(context.Background(), lt)
	return err
}

func GetLinkTraceList(page, pageSize int64) (lt []*LinkTrace, err error) {
	lt = []*LinkTrace{}

	err = global.MongoLinksTraceColl.Find(context.Background(),
		nil).Skip(pageSize * page).Limit(pageSize).All(&lt)
	if err == mongo.ErrNilDocument {
		return lt, nil
	}
	return
}

func GetLinkTraceListBySc(sc string, page, pageSize int64) (lt []*LinkTrace, err error) {
	lt = []*LinkTrace{}

	err = global.MongoLinksTraceColl.Find(context.Background(),
		bson.M{"sc": sc}).Skip(pageSize * page).Limit(pageSize).All(&lt)
	if err == mongo.ErrNilDocument {
		return lt, nil
	}
	return
}
