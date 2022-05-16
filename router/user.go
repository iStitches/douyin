package router

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {

	UserRouter := group.Group("user")

	UserRouter.GET("register", controller.Register)

	UserRouter.POST("login", controller.Login)
}
