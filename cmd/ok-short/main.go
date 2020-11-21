package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/model"
	"github.com/joeyscat/ok-short/internal/routers"
	"github.com/joeyscat/ok-short/pkg/logger"
	"github.com/joeyscat/ok-short/pkg/setting"
	"github.com/nats-io/nats.go"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// @title 短链接服务
// @version 1.0
// @description GoGo
// @termsOfService mm
func main() {
	initEnv()
	defer global.DBEngine.Close()
	defer global.Redis.Close()
	defer global.Nats.Close()

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

func initEnv() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("setupDBEngine err: %v", err)
	}

	err = setupRedis()
	if err != nil {
		log.Fatalf("setupRedis err: %v", err)
	}

	err = setupNats()
	if err != nil {
		log.Fatalf("setupNats err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("setupLogger err: %v", err)
	}
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Nats", &global.NatsSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	if global.ServerSetting.RunMode == "debug" {
		log.Printf("ServerSetting: %v", global.ServerSetting)
		log.Printf("AppSetting: %v", global.AppSetting)
		log.Printf("DatabaseSetting: %v", global.DatabaseSetting)
		log.Printf("Redis: %v", global.RedisSetting)
		log.Printf("Nats: %v", global.NatsSetting)
	}

	return nil
}

func setupLogger() error {
	var w io.Writer
	if global.ServerSetting.RunMode == "debug" {
		w = os.Stdout
	} else {
		w = &lumberjack.Logger{
			Filename: global.AppSetting.LogSavePath + "/" +
				global.AppSetting.LogFileName +
				global.AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}
	}

	global.Logger = logger.NewLogger(w, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupRedis() error {
	global.Redis = model.NewRedis(global.RedisSetting)
	if global.Redis == nil {
		return fmt.Errorf("cnnnect redis failed")
	}

	return nil
}

func setupNats() error {
	var err error
	global.Nats, err = nats.Connect(global.NatsSetting.Url)
	if err != nil {
		panic(err)
	}

	return nil
}
