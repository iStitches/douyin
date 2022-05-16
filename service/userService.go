package service

import (
	"douyin/dao"
	"douyin/global"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}


type UserLoginResponse struct {
	utils.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	utils.Response
	UserInfo UserInfo `json:"user"`
}

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]UserInfo{}

type UserService struct {
}

func NewUserService() *UserService{
	return &UserService{}
}


func (*UserService) Register(c *gin.Context, username string, password string) (interface{}, error) {
	token, err := utils.CreateToken(c, username)
	if err != nil {
		return UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "create token error"},
		}, err
	}

	// username exists
	if _, err := dao.NewUserDaoInstance().QueryUserByName(username); err!=nil {
		return UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "username already exist"},
		}, err
	}
	// Generate a snowflake ID.
	nodeID := global.SnowNode.Generate()
	id := nodeID.Int64()

	if err := dao.NewUserDaoInstance().CreateUser(id, username, password); err!=nil {
		return UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "User create error"},
		}, err
	}

	return UserLoginResponse{
		Response: utils.Response{StatusCode: 0},
		UserId:   id,
		Token:    token,
	}, nil
}

func (*UserService) Login(c *gin.Context, username string, password string) (interface{}, error) {
	token, err := utils.CreateToken(c, username)
	if err != nil {
		return UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "create token error"},
		}, err
	}

	user, err := dao.NewUserDaoInstance().QueryUserByName(username)
	if err!=nil {
		return UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		}, err
	}

	if user.Password != password {
		return UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "password indirect"},
		}, err
	}

	return UserLoginResponse{
		Response: utils.Response{StatusCode: 0},
		UserId:   user.ID,
		Token:    token,
	}, nil
}

func (*UserService) UserInfo(c *gin.Context) (interface{}, error) {
	token := c.Query("token")

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}