package initialize

import (
	"douyin/router"
	"github.com/gin-gonic/gin"
)

/**
  统一路由处理
*/
func Routers() *gin.Engine {
	Router := gin.Default()
	// 固定路由前缀
	ApiGroup := Router.Group("/douyin/")
	// 用户业务路由
	router.UserRouter(ApiGroup)
	return Router
}
