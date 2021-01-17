package service

import (
	"context"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestService_CreateLinkTrace(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set("X-Real-IP", "10.10.10.10")
	req.Header.Set("User-Agent", "ok-short-test")
	req.Header.Set("Cookie", "ok-short-test")
	req.UserAgent()
	c := e.NewContext(req, httptest.NewRecorder())

	type fields struct {
		ctx context.Context
	}
	type args struct {
		sc  string
		url string
		e   echo.Context
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
				sc:  "MjJe",
				url: "https://echo.labstack.com/guide/testing",
				e:   c,
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
			if err := svc.CreateLinkTrace(tt.args.sc, tt.args.url, tt.args.e); (err != nil) != tt.wantErr {
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
