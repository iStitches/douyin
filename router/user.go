package router

import (
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("user")
	{
		UserRouter.Use(middlewares.MidCors)
		// 用户列表
		UserRouter.GET("list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		// 登录注册
		UserRouter.POST("login", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
