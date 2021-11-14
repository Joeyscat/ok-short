package app

import (
	"github.com/joeyscat/ok-short/pkg/convert"
	"github.com/labstack/echo/v4"
)

var (
	defaultSPageSize = 20
	maxPageSize      = 40
)

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func GetPage(c echo.Context) int {
	page := convert.StrTo(c.QueryParam("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c echo.Context) int {
	pageSize := convert.StrTo(c.QueryParam("page_size")).MustInt()
	if pageSize <= 0 {
		return defaultSPageSize
	}
	if pageSize > maxPageSize {
		return maxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
