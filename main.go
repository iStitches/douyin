package main

import (
	"douyin/global"
	"douyin/initialize"
	"fmt"
)

func main() {
	//1. 初始化yaml配置
	initialize.InitConfig()
	//2. 启动logger
	initialize.InitLogger()
	//2. 数据库连接
	initialize.InitDB()

	initialize.InitSnowNode(1)
	//3. 初始化routers
	Router := initialize.Routers()
	//4. 启动服务

	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		fmt.Errorf("the error is %x", err)
	}
}
