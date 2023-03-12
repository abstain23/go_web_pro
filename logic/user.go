package logic

import (
	"errors"
	"gin-project/dao/mysql"
	"gin-project/models"
	"gin-project/pkg/jwt"
	"gin-project/pkg/snowflake"
)

var (
	ErrorUserExist = errors.New("用户已经存在")
)

func Register(params *models.ParamsRegister) (err error) {
	// 1. 判断用户存不存在
	var exist bool
	exist, err = mysql.CheckUserExist(params.Username)
	if err != nil {
		return err
	}
	if exist {
		return ErrorUserExist
	}
	// 2. 生成UID
	uID := snowflake.GenID()
	// 3. 保存进数据库
	user := &models.User{
		UserID:     uID,
		Username:   params.Username,
		Password:   params.Password,
		RePassword: params.RePassword,
		Email:      params.Email,
		Gender:     params.Gender,
	}
	err = mysql.InsertUser(user)

	return
}

func Login(params *models.ParamsLogin) (tokenString, rTokenString string, err error) {
	user := &models.User{
		Username: params.Username,
		Password: params.Password,
	}
	err = mysql.Login(user)
	if err != nil {
		return "", "", err
	}
	// 生成token
	return jwt.GenToken(user.UserID, params.Username)

}
