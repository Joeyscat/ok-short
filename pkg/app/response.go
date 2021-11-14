package app

import (
	"net/http"

	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
)

func Response(ctx echo.Context, data interface{}) error {
	if data == nil {
		data = map[string]interface{}{}
	}
	return ctx.JSON(http.StatusOK, data)
}

func ResponseList(ctx echo.Context, list interface{}, totalRows int) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
		"pager": Pager{
			Page:      GetPage(ctx),
			PageSize:  GetPageSize(ctx),
			TotalRows: totalRows,
		},
	})
}

func ErrorResponse(ctx echo.Context, err *errcode.Error) error {
	response := map[string]interface{}{"code": err.Code, "msg": err.Msg}
	details := err.Details
	if len(details) > 0 {
		response["details"] = details
	}
	return ctx.JSON(err.StatusCode(), response)
}
