package router

import (
	"douyin/controller"
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func PublishRouter(group *gin.RouterGroup)  {
	publishRouter := group.Group("publish")
	publishRouter.Use(middlewares.JWTAuth())
	publishRouter.GET("list", middlewares.JWTAuth(), controller.PublishList)

	publishRouter.POST("action", middlewares.JWTAuth(), controller.Publish)

}
