package initialize

import (
	"douyin/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"runtime/debug"
)

/**
  数据库初始化、启动
*/
func InitMysqlDB() error {
	// 启动失败错误处理，错误处理还未统一
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("InitMysql:%v", err))
			fmt.Println(fmt.Errorf("%s", string(debug.Stack())))
		}
	}()
	mysqlInfo := global.Settings.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.DB = db
	return nil
}
