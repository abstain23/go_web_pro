package controllers

import (
	"errors"
	"gin-project/middleware"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// 获取当前用户登录id
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middleware.ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
