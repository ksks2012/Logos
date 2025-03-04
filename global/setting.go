package global

import (
	"github.com/logos/pkg/logger"
	"github.com/logos/pkg/setting"
)

var (
	// AppSetting is the global app setting
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	SaveLoadSetting *setting.SaveLoadSettingS
	// CARSetting is the global character attribute range setting
	CARSetting *setting.CARSettingS
	ParSetting *setting.PARSettingS
	// WorldAttributeSetting is the global world setting
	WorldAttributeSetting *setting.WorldSettingS
	// GlobalAttributeSetting is the global setting
	GlobalAttributeSetting *setting.GlobalSettingS
	// Logger is the global logger
	Logger *logger.Logger
)
