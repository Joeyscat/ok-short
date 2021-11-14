package internal

import (
	"context"
	"net/http"
	"time"

	v1 "github.com/joeyscat/ok-short/internal/controller/v1"
	"github.com/joeyscat/ok-short/internal/global"
	"github.com/joeyscat/ok-short/internal/store"
	"github.com/joeyscat/ok-short/internal/store/mongo"
	"github.com/joeyscat/ok-short/pkg/app"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/joeyscat/ok-short/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() Server {

	// 初始化数据库连接
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	mongoStore, err := mongo.GetMongoFactoryOr(ctx, options.Client().ApplyURI(global.MongoDBSetting.URI))
	if err != nil {
		panic(err)
	}

	// 初始化路由
	e := echo.New()

	installMiddlewares(e)

	installRouters(e, mongoStore)

	return Server{echo: e}
}

func (s *Server) Start() error {

	return s.echo.Start(":" + global.AppSetting.HttpPort)
}

func installMiddlewares(e *echo.Echo) {
	e.Validator = app.NewValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	allowOrigins := global.AppSetting.AllowOrigins
	if len(allowOrigins) == 0 {
		panic("allowOrigins should not be empty")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}

func installRouters(e *echo.Echo, factory store.Factory) {
	apiV2 := e.Group("/api/v2")
	link := v1.NewLinkController(factory)
	linkTrace := v1.NewLinkTraceController(factory)
	e.GET("/:sc", link.Redirect)
	{
		// 创建短链接
		apiV2.POST("/links", link.Shorten)
		apiV2.GET("/links/:sc", link.Get)
		apiV2.GET("/links", link.List)

		apiV2.GET("/link-trace/:sc", linkTrace.Get)
		apiV2.GET("/link-trace", linkTrace.List)
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
