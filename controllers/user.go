package controllers

import (
	"gin-project/logic"
	"gin-project/models"
	innerValidator "gin-project/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {

	var p = new(models.ParamsRegister)
	// 1. 参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": innerValidator.RemoveTopStruct(errors.Translate(innerValidator.Trans)),
		})
		return

	}

	// 2. 业务处理
	logic.Register(p)

	// 3. 返回响应
	c.JSON(http.StatusOK, "ok")
}
