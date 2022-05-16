package dao

import (
	"douyin/global"
	"douyin/models"
	"gorm.io/gorm"
	"sync"
)

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

//单例模式，参照字节第二节课的https://github.com/Moonlight-Zhao/go-project-example进行的设计
//暂时还没研究这种设计模式是否合理
func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) CreateUser(id int64, username string, password string) error {
	user := models.User{ID:id, UserName: username, Password: password}
	result := global.DB.Create(&user)
	return result.Error
}

func (*UserDao) QueryUserByName(name string) (*models.User, error) {
	var user models.User
	err := global.DB.Where("username = ?", name).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		global.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserById(id int64) (*models.User, error) {
	var user models.User
	err := global.DB.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		global.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

//
//func (*UserDao) MQueryUserById(ids []int64) (map[int64]*models.User, error) {
//	var users []*models.User
//	err := global.DB.Where("id in (?)", ids).Find(&users).Error
//	if err != nil {
//		global.Logger.Error("batch find user by id err:" + err.Error())
//		return nil, err
//	}
//	userMap := make(map[int64]*models.User)
//	for _, user := range users {
//		userMap[user.ID] = user
//	}
//	return userMap, nil
//}