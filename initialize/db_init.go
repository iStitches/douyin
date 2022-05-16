package initialize

import (
	"douyin/global"
	"douyin/models"
	"douyin/utils"
)

func InitDB() error {

	var err error
	global.DB, err = utils.InitMysqlDB()
	if err !=nil {
		global.Logger.Error("find post by id err:" + err.Error())
	}
	// 迁移
	var models = []interface{}{&models.User{}, &models.Video{}}
	global.DB.AutoMigrate(models...)

	return nil
}
