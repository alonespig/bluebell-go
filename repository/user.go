package repository

import (
	"bluebell/global"
	"bluebell/model"
)

func CreateUser(user *model.User) error {
	return global.DB.Create(user).Error
}

// CheckUserNameUnique 检查用户名是否唯一
func CheckUserNameUnique(username string) bool {
	var count int64
	if err := global.DB.Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false
	}
	return count == 0
}

// GetUserByID 根据ID获取用户
func GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
