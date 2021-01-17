package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/routers"
	"log"
	"net/http"
)

// @title 短链接服务
// @version 1.0
// @description GoGo
// @termsOfService mm
func main() {
	global.InitEnv()
	defer global.Redis.Close()

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	defer s.Close()
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("server err: %v", err)
	}
}
