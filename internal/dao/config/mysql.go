package config

import (
	bsinterfaces "github.com/common_structure/interfaces"
	mysqlstorage "github.com/common_structure/internal/dao/dbversion/mysql"
	"github.com/common_structure/pkg/setting"
)

func setupMySQLStorageEngine(databaseSetting *setting.DatabaseSettingS) (storage bsinterfaces.StorageEngine, err error) {
	if len(databaseSetting.Host) == 0 {
		return
	}
	return mysqlstorage.NewMySQLStorageEngine(databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host[0], databaseSetting.SocketPath, databaseSetting.DBName)
}

func setupMySQLRoundRobinStorageEngine(databaseSetting *setting.DatabaseSettingS) (storage bsinterfaces.StorageEngine, err error) {
	return mysqlstorage.NewMySQLRoundRobinStorageEngine(databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host, databaseSetting.SocketPath, databaseSetting.DBName)
}
