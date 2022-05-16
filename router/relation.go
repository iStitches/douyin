package router

import (
	"douyin/controller"
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func RelationRouter(group *gin.RouterGroup)  {
	relationRouter := group.Group("relation")
	relationRouter.Use(middlewares.JWTAuth())
	relationRouter.GET("follow/list", controller.FollowList)
	relationRouter.GET("follower/list", controller.FollowerList)

	relationRouter.POST("action", controller.RelationAction)

}
