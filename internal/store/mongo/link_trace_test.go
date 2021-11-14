package mongo

import (
	"context"
	"testing"

	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/store"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_linkTraces_CreateLinkTrace(t *testing.T) {
	ctx := context.Background()
	factory, err := GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.LinkTraceStore
	}
	type args struct {
		ctx context.Context
		lt  *model.LinkTrace
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"OK",
			fields{factory.LinkTraces()},
			args{ctx, &model.LinkTrace{
				Sc:  "ok",
				URL: "www",
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.store
			if err := s.CreateLinkTrace(tt.args.ctx, tt.args.lt); (err != nil) != tt.wantErr {
				t.Errorf("linkTraces.CreateLinkTrace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_linkTraces_ListLinkTrace(t *testing.T) {
	ctx := context.Background()
	factory, err := GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.LinkTraceStore
	}
	type args struct {
		ctx      context.Context
		page     int64
		pageSize int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantLt  []*model.LinkTrace
		wantErr bool
	}{
		{
			"OK",
			fields{factory.LinkTraces()},
			args{ctx, 1, 10},
			[]*model.LinkTrace{{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.store
			gotLt, err := s.ListLinkTrace(tt.args.ctx, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("linkTraces.ListLinkTrace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.wantLt) > 0 {
				assert.NotEmpty(t, gotLt)

				for _, v := range gotLt {
					t.Logf("%+v", v)
				}
			} else {
				assert.Empty(t, gotLt)
			}
		})
	}
}

func Test_linkTraces_ListLinkTraceBySc(t *testing.T) {
	ctx := context.Background()
	factory, err := GetMongoFactoryOr(ctx, options.Client().ApplyURI(test_mongodb_uri))
	assert.Nil(t, err)

	type fields struct {
		store store.LinkTraceStore
	}
	type args struct {
		ctx      context.Context
		sc       string
		page     int64
		pageSize int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantLt  []*model.LinkTrace
		wantErr bool
	}{
		{
			"OK",
			fields{factory.LinkTraces()},
			args{ctx, "", 1, 10},
			nil,
			true,
		},
		{
			"OK",
			fields{factory.LinkTraces()},
			args{ctx, "ok", 1, 10},
			[]*model.LinkTrace{{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.store
			gotLt, err := s.ListLinkTraceBySc(tt.args.ctx, tt.args.sc, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("linkTraces.ListLinkTraceBySc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.wantLt) > 0 {
				assert.NotEmpty(t, gotLt)

				for _, v := range gotLt {
					t.Logf("%+v", v)
				}
			} else {
				assert.Empty(t, gotLt)
			}
		})
	}
}
