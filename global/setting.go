package global

import (
	"github.com/joeyscat/ok-short/pkg/logger"
	"github.com/joeyscat/ok-short/pkg/setting"
)

var (
	AppSetting     *setting.AppSettingS
	MongoDBSetting *setting.MongoDBSettingS
	RedisSetting   *setting.RedisSettingS
	Logger         *logger.Logger
)
