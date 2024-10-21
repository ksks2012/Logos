package global

import (
	"github.com/logos/pkg/logger"
	"github.com/logos/pkg/setting"
)

var (
	// AppSetting is the global app setting
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	// CARSetting is the global character attribute range setting
	CARSetting *setting.CARSettingS
	ParSetting *setting.PARSettingS
	// Logger is the global logger
	Logger *logger.Logger
)
