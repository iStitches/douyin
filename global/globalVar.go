package global

import (
	"douyin/config"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	DB       *gorm.DB
)
