package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/pkg/convert"
)

var (
	defaultSPageSize = 20
	maxPageSize      = 40
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
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
