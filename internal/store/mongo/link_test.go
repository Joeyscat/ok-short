package mongo

import (
	"context"
	"reflect"
	"testing"

	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/store"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	test_mongodb_uri = "mongodb://ok-short:123456@app.io:27017/ok-short"
)

func Test_links_ListLinks(t *testing.T) {
	ctx := context.Background()
	factory, err := GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.LinkStore
	}
	type args struct {
		ctx      context.Context
		page     int64
		pageSize int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList []*model.Link
		wantErr  bool
	}{
		{
			"OK",
			fields{factory.Links()},
			args{ctx, 1, 10},
			[]*model.Link{{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.store
			gotList, err := s.ListLinks(tt.args.ctx, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("links.ListLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(tt.wantList) > 0 {
				assert.NotEmpty(t, gotList)

				for _, v := range gotList {
					t.Logf("%+v", v)
				}
			} else {
				assert.Empty(t, gotList)
			}
		})
	}
}

func Test_links_CreateLink(t *testing.T) {
	ctx := context.Background()
	factory, err := GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.LinkStore
	}
	type args struct {
		ctx  context.Context
		link *model.Link
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"OK",
			fields{factory.Links()},
			args{ctx, &model.Link{
				Sc: "ok", Status: "xx", OriginURL: "www", Exp: 1000,
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.store
			if err := s.CreateLink(tt.args.ctx, tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("links.CreateLink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_links_GetLinkBySc(t *testing.T) {
	ctx := context.Background()
	factory, err := GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.LinkStore
	}
	type args struct {
		ctx context.Context
		sc  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantL   *model.Link
		wantErr bool
	}{
		// {
		// 	"OK",
		// 	fields{factory.Links()},
		// 	args{ctx, ""},
		// 	&model.Link{},
		// 	false,
		// },
		{
			"OK",
			fields{factory.Links()},
			args{ctx, "x"},
			nil,
			false,
		},
		// {
		// 	"OK",
		// 	fields{factory.Links()},
		// 	args{ctx, "ok"},
		// 	&model.Link{Sc: "ok", Status: "xx", OriginURL: "www", Exp: 1000},
		// 	false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.store
			gotL, err := s.GetLinkBySc(tt.args.ctx, tt.args.sc)
			if (err != nil) != tt.wantErr {
				t.Errorf("links.GetLinkBySc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("links.GetLinkBySc() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
