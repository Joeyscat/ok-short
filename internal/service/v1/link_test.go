package v1

import (
	"context"
	"testing"

	"github.com/joeyscat/ok-short/internal/store"
	"github.com/joeyscat/ok-short/internal/store/mongo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	test_mongodb_uri = "mongodb://ok-short:123456@app.io:27017/ok-short"
)

func Test_linkService_CreateLink(t *testing.T) {
	ctx := context.Background()
	factory, err := mongo.GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.Factory
	}
	type args struct {
		ctx   context.Context
		param *CreateLinkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"OK",
			fields{store: factory},
			args{},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &linkService{
				store: tt.fields.store,
			}
			got, err := s.CreateLink(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("linkService.CreateLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("linkService.CreateLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkService_UnShorten(t *testing.T) {
	ctx := context.Background()
	factory, err := mongo.GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.Factory
	}
	type args struct {
		ctx   context.Context
		param *RedirectLinkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"OK",
			fields{store: factory},
			args{context.Background(), &RedirectLinkRequest{Sc: "not_exists"}},
			"",
			true,
		},
		{
			"OK",
			fields{store: factory},
			args{context.Background(), &RedirectLinkRequest{Sc: "ok"}},
			"www",
			false,
		},
		{
			"OK",
			fields{store: factory},
			args{context.Background(), &RedirectLinkRequest{Sc: ""}},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &linkService{
				store: tt.fields.store,
			}
			got, err := s.UnShorten(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("linkService.UnShorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("linkService.UnShorten() = %v, want %v", got, tt.want)
			}
		})
	}
}
