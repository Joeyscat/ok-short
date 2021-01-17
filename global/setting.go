package global

import (
	"github.com/joeyscat/ok-short/pkg/logger"
	"github.com/joeyscat/ok-short/pkg/setting"
)

var (
	ServerSetting  *setting.ServerSettingS
	AppSetting     *setting.AppSettingS
	MongoDBSetting *setting.MongoDBSettingS
	RedisSetting   *setting.RedisSettingS
	Logger         *logger.Logger
	JWTSetting     *setting.JWTSettingS
)
