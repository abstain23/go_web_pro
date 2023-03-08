package logic

import (
	"errors"
	"gin-project/dao/mysql"
	"gin-project/models"
	"gin-project/pkg/snowflake"
)

func Register(params *models.ParamsRegister) (err error) {
	// 1. 判断用户存不存在
	var exist bool
	exist, err = mysql.CheckUserExist(params.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("user exist")
	}
	// 2. 生成UID
	uID := snowflake.GenID()
	// 3. 保存进数据库
	user := &models.User{
		UserID: uID,
		ParamsRegister: models.ParamsRegister{
			Username:   params.Username,
			Password:   params.Password,
			RePassword: params.RePassword,
			Email:      params.Email,
			Gender:     params.Gender,
		},
	}
	mysql.InsertUser(user)

	return
}
