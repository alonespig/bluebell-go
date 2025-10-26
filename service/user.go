package service

import (
	"bluebell/froms"
	"bluebell/jwt"
	"bluebell/model"
	"bluebell/pkg/snowflake"
	"bluebell/repository"
	"errors"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) error {
	if !repository.CheckUserNameUnique(username) {
		return errors.New("username already exists")
	}

	//生成哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Error("failed to generate password hash", err)
		return err
	}
	user := model.User{
		UserID:   snowflake.GenID(),
		Username: username,
		Password: string(hash),
	}
	return repository.CreateUser(&user)
}

// CheckPassword 检查密码是否正确
func CheckPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// Login 登录
func Login(username, password string) (*froms.LoginResponse, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		zap.S().Error("failed to get user by username", err)
		return nil, err
	}
	if !CheckPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	token, err := jwt.GenerateToken(user.UserID, user.Username)
	if err != nil {
		zap.S().Error("failed to generate token", err)
		return nil, err
	}

	return &froms.LoginResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Token:    token,
	}, nil
}
