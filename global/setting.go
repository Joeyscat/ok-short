package global

import (
	"github.com/joeyscat/ok-short/pkg/logger"
	"github.com/joeyscat/ok-short/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
