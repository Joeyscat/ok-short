package v2

import (
	_ "github.com/joeyscat/ok-short/docs"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/middleware"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
	echoMiddle "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Validator = app.NewValidator()
	e.HTTPErrorHandler = httpErrorHandler

	e.Use(echoMiddle.Logger())
	e.Use(middleware.Recovery)
	e.Use(echoMiddle.CORSWithConfig(echoMiddle.CORSConfig{
		AllowOrigins: []string{"http://oook.fun", "http://u.oook.fun"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	link := NewLink()
	linkTrace := NewLinkTrace()
	e.GET("/:sc", link.Redirect)
	{
		apiV2 := e.Group("/api/v2")
		//apiV2.GET("/auth", GetAuth)
		apiV2.GET("/login", Login)

		apiV2.Use(echoMiddle.BasicAuth(Auth))

		// 创建短链接
		apiV2.POST("/links", link.Shorten)
		apiV2.GET("/links/:sc", link.Get)
		apiV2.GET("/links", link.List)

		apiV2.GET("/link-trace/:sc", linkTrace.Get)
		apiV2.GET("/link-trace", linkTrace.List)
	}

	return e
}

func httpErrorHandler(err error, c echo.Context) {
	var (
		code    = errcode.ServerError.Code()
		msg     = errcode.ServerError.Msg()
		details = errcode.ServerError.Details()
	)

	if e, ok := err.(*errcode.Error); ok {
		code = e.Code()
		msg = e.Msg()
		details = e.Details()
	} else if global.ServerSetting.RunMode == "debug" {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			//response := app.NewResponse(c)
			//if err := response.ToErrorResponse(errcode.ServerError);err!= nil {
			//	c.Logger().Error(err)
			//}

			err := c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    code,
				"msg":     msg,
				"details": details,
			})
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
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
