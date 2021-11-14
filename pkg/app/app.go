package app

import (
	"net/http"

	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Ctx echo.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx echo.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) error {
	if data == nil {
		data = map[string]interface{}{}
	}
	return r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) error {
	return r.Ctx.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) error {
	response := map[string]interface{}{"code": err.Code, "msg": err.Msg}
	details := err.Details
	if len(details) > 0 {
		response["details"] = details
	}
	return r.Ctx.JSON(err.StatusCode(), response)
}
