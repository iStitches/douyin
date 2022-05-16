package controller

import (
	"douyin/global"
	"douyin/service"
	"douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)





var userService = service.NewUserService()


func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	respose, err := userService.Register(c,username,password)
	if err != nil {
		fmt.Println("register error:",err.Error())
	}
	c.JSON(http.StatusOK,respose)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	respose, err := userService.Login(c,username,password)
	if err != nil {
		fmt.Println("login error:",err.Error())
	}
	c.JSON(http.StatusOK,respose)
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	//userService.UserInfo(c)
}