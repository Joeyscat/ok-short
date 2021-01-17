package global

import (
	"errors"
	"fmt"
	"github.com/joeyscat/ok-short/pkg/logger"
	"github.com/joeyscat/ok-short/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"time"
)

func InitEnv() {
	err := SetupSetting()
	if err != nil {
		log.Fatalf("SetupSetting err: %v", err)
	}

	err = SetupMongoDB()
	if err != nil {
		log.Fatalf("SetupMongoDB err: %v", err)
	}

	err = SetupRedis()
	if err != nil {
		log.Fatalf("SetupRedis err: %v", err)
	}

	err = SetupLogger()
	if err != nil {
		log.Fatalf("SetupLogger err: %v", err)
	}
}

func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	if ServerSetting == nil {
		return errors.New("Configuration not found: ServerSetting. ")
	}
	err = s.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}
	if AppSetting == nil {
		return errors.New("Configuration not found: AppSetting. ")
	}
	err = s.ReadSection("MongoDB", &MongoDBSetting)
	if err != nil {
		return err
	}
	if MongoDBSetting == nil {
		return errors.New("Configuration not found: MongoDBSetting. ")
	}
	err = s.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}
	if RedisSetting == nil {
		return errors.New("Configuration not found: RedisSetting. ")
	}
	err = s.ReadSection("JWT", &JWTSetting)
	if err != nil {
		return err
	}
	if JWTSetting == nil {
		return errors.New("Configuration not found: JWTSetting. ")
	}

	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	JWTSetting.Expire *= time.Second

	if ServerSetting.RunMode == "debug" {
		log.Printf("ServerSetting: %v", ServerSetting)
		log.Printf("AppSetting: %v", AppSetting)
		log.Printf("Redis: %v", RedisSetting)
	}

	return nil
}

func SetupLogger() error {
	var w io.Writer
	if ServerSetting.RunMode == "debug" {
		w = os.Stdout
	} else {
		w = &lumberjack.Logger{
			Filename: AppSetting.LogSavePath + "/" +
				AppSetting.LogFileName +
				AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}
	}

	Logger = logger.NewLogger(w, "", log.LstdFlags).WithCaller(2)

	return nil
}

func SetupMongoDB() error {
	var err error
	mongoDB, err := NewMongoDB(MongoDBSetting)
	if err != nil {
		return err
	}
	MongoOkShortDB = mongoDB.Database("ok-short")
	MongoLinksColl = MongoOkShortDB.Collection("links")
	MongoLinksTraceColl = MongoOkShortDB.Collection("links_traces")
	MongoAuthsColl = MongoOkShortDB.Collection("auths")

	return nil
}

func SetupRedis() error {
	Redis = NewRedis(RedisSetting)
	if Redis == nil {
		return fmt.Errorf("cnnnect redis failed")
	}

	return nil
}
