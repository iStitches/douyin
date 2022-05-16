package initialize

import (
	"douyin/global"
	"go.uber.org/zap"
)

func InitLogger() error {
	var err error
	global.Logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	return nil
}