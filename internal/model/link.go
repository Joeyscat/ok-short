package model

import (
	"context"
	"github.com/joeyscat/ok-short/internal/global"
	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Link struct {
	field.DefaultField `bson:",inline"`

	Sc        string `bson:"sc" json:"sc"`                 // 短链代码
	Status    string `bson:"status" json:"status"`         // 短链状态
	OriginURL string `bson:"origin_url" json:"origin_url"` // 原始链接
	Exp       uint32 `bson:"exp" json:"exp"`               // 过期时间
}

func CreateLink(sc string, originalURL string, exp uint32) (*Link, error) {
	l := &Link{
		Sc:        sc,
		OriginURL: originalURL,
		Exp:       exp,
	}
	_, err := global.MongoLinksColl.InsertOne(context.Background(), l)
	return l, err
}

func GetLinkDetailBySc(sc string) (l *Link, err error) {
	l = &Link{}
	err = global.MongoLinksColl.Find(context.Background(),
		bson.M{"sc": sc}).One(l)
	if err == mongo.ErrNilDocument {
		return l, nil
	}
	return
}

func GetLinkBySc(sc string) (link string, err error) {
	l := &Link{}
	// 覆盖索引
	err = global.MongoLinksColl.Find(context.Background(),
		bson.M{"sc": sc}).Select(bson.M{"_id": 0, "origin_url": 1}).One(l)
	link = l.OriginURL

	if err == mongo.ErrNoDocuments {
		return "", nil
	}

	return
}

func GetLinkList(page, pageSize int64) (list []*Link, err error) {
	list = []*Link{}

	err = global.MongoLinksColl.Find(context.Background(),
		bson.M{}).Skip(pageSize * page).Limit(pageSize).All(&list)
	if err == mongo.ErrNilDocument {
		return list, nil
	}

	return
}

func CreateIndex() error {
	return global.MongoLinksColl.CreateIndexes(
		context.Background(),
		[]options.IndexModel{{Key: []string{"sc", "origin_url"}}},
	)
}
