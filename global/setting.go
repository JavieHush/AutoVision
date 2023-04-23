package global

import (
	"AutoVision/pkg/logger"
	"AutoVision/pkg/setting"
)

// link configs with applications

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
