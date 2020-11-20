package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joeyscat/ok-short/docs"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/middleware"
	"github.com/joeyscat/ok-short/internal/routers/api"
	v1 "github.com/joeyscat/ok-short/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(Cors())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	//r.Use(middleware.RateLimiter(methodLimiters))
	//r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	link := v1.NewLink()
	linkTrace := v1.NewLinkTrace()
	//upload := api.NewUpload()
	//r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload/file", upload.UploadFile)
	//r.POST("/auth", api.GetAuth)
	//r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.GET("/auth", api.GetAuth)
	// 短链接跳转
	// 这里 /a 是因为gin不能在跟路径进行匹配
	r.GET("/a/:sc", link.Redirect)
	apiV1 := r.Group("/api/v1")
	apiV1.Use() //middleware.JWT()
	{
		// 创建短链接
		apiV1.POST("/links", link.Shorten)
		apiV1.GET("/links/:sc", link.Get)
		apiV1.GET("/links", link.List)

		apiV1.GET("/link-trace/:sc", linkTrace.Get)
		apiV1.GET("/link-trace", linkTrace.List)
	}

	return r
}

// 用于接受Token的Header
const TokenHeaderName = "OK-Short-Token"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//接收客户端发送的origin （重要！）
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//允许跨域设置可以返回其他子段，可以自定义字段
		c.Header("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, cross-request-open-sign, "+TokenHeaderName)
		// 允许浏览器（客户端）可以解析的头部 （重要）
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, OK-Short-Token")
		//设置缓存时间
		c.Header("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		c.Header("Access-Control-Allow-Credentials", "true")

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		c.Next()
	}
}
