package global

import (
	"github.com/mneumi/gin-boilerplate/pkg/logger"
	"github.com/mneumi/gin-boilerplate/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
