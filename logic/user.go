package logic

import (
	"gin-project/dao/mysql"
	"gin-project/pkg/snowflake"
)

func Register() {
	// 1. 判断用户存不存在
	mysql.QueryUserByUsername()
	// 2. 生成UID
	snowflake.GenID()
	// 3. 保存进数据库
	mysql.InsertUser()
}
