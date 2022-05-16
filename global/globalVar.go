package global

import (
	"douyin/config"
	"douyin/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings    config.ServerConfig
	DB          *gorm.DB
    Logger      *zap.Logger
	SnowNode  *utils.Node
)

