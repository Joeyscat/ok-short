package model

import (
    "context"
    "github.com/go-redis/redis"
    global2 "github.com/joeyscat/ok-short/internal/global"
    "github.com/joeyscat/ok-short/pkg/setting"
    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "reflect"
    "testing"
)

func TestNewMongoDB(t *testing.T) {
	type args struct {
		s *setting.MongoDBSettingS
	}
	tests := []struct {
		name    string
		args    args
		want    *mongo.Client
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				s: &setting.MongoDBSettingS{
					Addr:     []string{"192.168.50.119:27017"},
					User:     "ok-short_rw",
					Password: "123456",
					AuthDB:   "db_ok_short",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	type UserInfo struct {
		Name   string `bson:"name"`
		Age    uint16 `bson:"age"`
		Weight uint32 `bson:"weight"`
	}

	var userInfo = UserInfo{
		Name:   "xm",
		Age:    7,
		Weight: 40,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := global2.NewMongoDB(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			t.Logf("%v", got)

			db := got.Database(tt.args.s.AuthDB)
			t.Logf("%v", &db)
			coll := db.Collection("test-coll")

			ctx := context.Background()
			_, err = coll.InsertOne(ctx, userInfo)
			assert.Nil(t, err)
			one := UserInfo{}
			err = coll.Find(ctx, bson.M{"name": userInfo.Name}).One(&one)
			assert.Nil(t, err)
			t.Logf("%v", one)
		})
	}
}

func TestNewRedis(t *testing.T) {
	type args struct {
		s *setting.RedisSettingS
	}
	tests := []struct {
		name string
		args args
		want *redis.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := global2.NewRedis(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
