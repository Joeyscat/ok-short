package model

import (
	"github.com/go-redis/redis"
	"github.com/joeyscat/ok-short/internal/global"
	"github.com/joeyscat/ok-short/pkg/setting"
	"github.com/stretchr/testify/assert"
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := global.NewMongoDB(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			t.Logf("%v", got)
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
			if got := global.NewRedis(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
