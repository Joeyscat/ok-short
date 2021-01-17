package v2

import (
	_ "github.com/joeyscat/ok-short/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://oook.fun", "http://u.oook.fun"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	apiV2 := e.Group("/api/v2")
	link := NewLink()
	linkTrace := NewLinkTrace()
	e.GET("/:sc", link.Redirect)
	{
		apiV2.GET("/auth", GetAuth)

		// 创建短链接
		apiV2.POST("/links", link.Shorten)
		apiV2.GET("/links/:sc", link.Get)
		apiV2.GET("/links", link.List)

		apiV2.GET("/link-trace/:sc", linkTrace.Get)
		apiV2.GET("/link-trace", linkTrace.List)
	}

	return e
}

// 用于接受Token的Header
//const TokenHeaderName = "OK-Short-Token"

//func Cors() echo.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		//接收客户端发送的origin （重要！）
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		//允许跨域设置可以返回其他子段，可以自定义字段
//		c.Header("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, cross-request-open-sign, "+TokenHeaderName)
//		// 允许浏览器（客户端）可以解析的头部 （重要）
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, OK-Short-Token")
//		//设置缓存时间
//		c.Header("Access-Control-Max-Age", "172800")
//		//允许客户端传递校验信息比如 cookie (重要)
//		c.Header("Access-Control-Allow-Credentials", "true")
//
//		//允许类型校验
//		if method == "OPTIONS" {
//			c.JSON(http.StatusOK, "ok!")
//		}
//
//		c.Next()
//	}
//}
