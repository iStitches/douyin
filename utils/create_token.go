package utils

import (
	"douyin/middlewares"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateToken(c *gin.Context, UserName string) (string, error) {
	//生成token信息
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		UserName:    UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60*60*24*1, //token -->1天过期
			Issuer:    "douyin",
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	return token, err
}