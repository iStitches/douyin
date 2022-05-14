package main

import (
	"douyin/initialize"
	"fmt"
)

func main() {
	//1. 初始化yaml配置
	initialize.InitConfig()
	//2. 数据库连接
	initialize.InitMysqlDB()
	//3. 初始化routers
	Router := initialize.Routers()
	//4. 启动服务
	err := Router.Run()
	if err != nil {
		fmt.Errorf("the error is %x", err)
	}
}
