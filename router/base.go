package router

import (
	"douyin/controller"
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func BaseRouter(group *gin.RouterGroup) {
	// basic apis
	group.GET("feed", controller.Feed)
	group.GET("user", middlewares.JWTAuth(),controller.UserInfo)

}