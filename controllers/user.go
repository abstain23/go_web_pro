package controllers

import (
	"gin-project/logic"
	"gin-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {

	var p = new(models.ParamsRegister)
	// 1. 参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}

	// 2. 业务处理
	logic.Register()

	// 3. 返回响应
	c.JSON(http.StatusOK, "ok")
}
