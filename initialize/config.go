package initialize

import (
	"douyin/config"
	"douyin/global"
	"github.com/spf13/viper"
)

/**
  viper读取解析配置文件
*/
func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./setting-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	global.Settings = serverConfig
}
