package initialize

import (
	"douyin/global"
	"douyin/utils"
)

func InitSnowNode(node int64) error {
	var err error
	global.SnowNode, err = utils.NewNode(1)
	if err != nil {
		return err
	}
	return nil
}
