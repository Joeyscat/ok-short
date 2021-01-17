package service

import (
	"context"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 准备测试环境
func setup() {
	viper.AddConfigPath("D:\\dev\\code\\go\\ok-short\\configs\\")
	err := global.SetupSetting()
	if err != nil {
		panic(err)
	}
	err = global.SetupRedis()
	if err != nil {
		panic(err)
	}
	err = global.SetupMongoDB()
	if err != nil {
		panic(err)
	}
}

func TestService_CreateLink(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
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
			name: "创建短链接",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				&CreateLinkRequest{
					URL:                 "https://github.com/qiniu/qmgo/blob/master/README_ZH.md",
					ExpirationInMinutes: 0,
				},
			},
			wantErr: false,
		},
	}
	setup()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				ctx: tt.fields.ctx,
			}
			got, err := svc.CreateLink(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != "" && got != tt.want {
				t.Errorf("CreateLink() got = %v, want %v", got, tt.want)
			}
			t.Logf("got: %s", got)
		})
	}
}

func TestService_GetLink(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		param *GetLinkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Link
		wantErr bool
	}{
		{
			name: "查询短链接",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				&GetLinkRequest{Sc: "MJR"},
			},
			wantErr: false,
		},
	}
	setup()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				ctx: tt.fields.ctx,
			}
			got, err := svc.GetLink(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}

func TestService_GetLinkList(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		pager *app.Pager
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Link
		wantErr bool
	}{
		{
			name: "查询短链接列表",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				&app.Pager{},
			},
			wantErr: false,
		},
	}
	setup()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				ctx: tt.fields.ctx,
			}
			got, err := svc.GetLinkList(tt.args.pager)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinkList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			for _, l := range got {
				t.Logf("link: %+v", l)
			}
		})
	}
}

func TestService_UnShorten(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
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
			name: "短连接还原",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				&RedirectLinkRequest{Sc: "MJR"},
			},
			wantErr: false,
		},
	}
	setup()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				ctx: tt.fields.ctx,
			}
			got, err := svc.UnShorten(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnShorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("get: %v", got)
		})
	}
}

func Test_genId(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genId()
			if (err != nil) != tt.wantErr {
				t.Errorf("genId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("genId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndex(t *testing.T) {
	setup()

	model.CreateIndex()
}
