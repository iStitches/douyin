package router

import (
	"douyin/controller"
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func CommentRouter(group *gin.RouterGroup)  {
	publishRouter := group.Group("publish")
	publishRouter.Use(middlewares.JWTAuth())
	publishRouter.GET("list", middlewares.JWTAuth(), controller.FavoriteList)

	publishRouter.POST("action", middlewares.JWTAuth(), controller.FavoriteAction)

}


