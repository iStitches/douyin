package utils

import (
	"douyin/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
  数据库初始化、启动
*/
func InitMysqlDB() (db *gorm.DB, err error) {
	// 启动失败错误处理，错误处理还未统一
	mysqlInfo := global.Settings.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
