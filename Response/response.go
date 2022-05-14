package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response interface {
	code() int
	msg() interface{}
	data() interface{}
}

type UserRegister struct {
	Id         int64  `json:"user_id"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	StatusCode int32  `json:"code"`
}

func (user UserRegister) code() int {
	return int(user.StatusCode)
}
func (user UserRegister) msg() interface{} {
	return user.StatusMsg
}
func (user UserRegister) data() interface{} {
	res := make(map[string]interface{})
	res["id"] = user.Id
	res["token"] = user.Token
	return res
}

func Success(c *gin.Context, code int, msg interface{}, data interface{}, response Response) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": response.code(),
		"msg":  response.msg(),
		"data": response.data(),
	})
	return
}
