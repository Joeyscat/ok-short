package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateLinkTrace(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		sc  string
		url string
		c   *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "创建链接访问记录",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				sc:  "aaa",
				url: "https://www.google.com/search?newwindow=1&client=firefox-b-d&sxsrf=ALeKk002xTpPWf3ybePeApicIIRNFKjNtw%3A1610810423423&ei=NwQDYMCqGdrCz7sP36uHyAQ&q=mgo+find+options+&oq=mgo+find+options+&gs_lcp=CgZwc3ktYWIQA1DnSljda2C2cmgBcAB4AIAB7AKIAa0JkgEHMC43LjAuMZgBAKABAaoBB2d3cy13aXrAAQE&sclient=psy-ab&ved=0ahUKEwiAotq44KDuAhVa4XMBHd_VAUkQ4dUDCAw&uact=5",
				c:   nil,
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
			if err := svc.CreateLinkTrace(tt.args.sc, tt.args.url, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateLinkTrace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetLinkTrace(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		param *GetLinkTraceRequest
		pager *app.Pager
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.LinkTrace
		wantErr bool
	}{
		{
			name: "查询指定短链接访问记录",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				param: &GetLinkTraceRequest{},
				pager: &app.Pager{
					Page:     0,
					PageSize: 20,
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
			got, err := svc.GetLinkTrace(tt.args.param, tt.args.pager)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinkTrace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)

			for _, trace := range got {
				t.Logf("trace: %+v", trace)
			}
		})
	}
}

func TestService_GetLinkTraceList(t *testing.T) {
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
		want    []*model.LinkTrace
		wantErr bool
	}{
		{
			name: "查询短链接访问记录",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				pager: &app.Pager{
					Page:     0,
					PageSize: 20,
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
			got, err := svc.GetLinkTraceList(tt.args.pager)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinkTraceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)

			for _, trace := range got {
				t.Logf("trace: %+v", trace)
			}
		})
	}
}
