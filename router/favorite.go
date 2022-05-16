package router

import (
	"douyin/controller"
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func FavoriteRouter(group *gin.RouterGroup)  {
	favoriteRouter := group.Group("comment")
	favoriteRouter.Use(middlewares.JWTAuth())
	favoriteRouter.GET("list/", middlewares.JWTAuth(), controller.CommentList)

	favoriteRouter.POST("action/", middlewares.JWTAuth(), controller.CommentAction)

}

