package controllers

import (
	"gin-project/constants"
	"gin-project/logic"
	"gin-project/models"
	innerValidator "gin-project/pkg/validator"
	"gin-project/utils"

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
			utils.ResponseError(c, constants.CodeInvalidParam)
			return
		}

		utils.ResponseWithCustomMsg(
			c,
			constants.CodeInvalidParam,
			innerValidator.RemoveTopStruct(errors.Translate(innerValidator.Trans)),
			nil,
		)
		return

	}

	// 2. 业务处理
	err := logic.Register(p)
	if err != nil {
		utils.ResponseWithCustomMsg(c, constants.CodeServerBusy, err.Error(), nil)
		return
	}

	// 3. 返回响应
	utils.ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamsLogin)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseWithCustomMsg(c, constants.CodeInvalidParam, err.Error(), nil)
			return
		}
		utils.ResponseWithCustomMsg(
			c,
			constants.CodeInvalidParam,
			innerValidator.RemoveTopStruct(errors.Translate(innerValidator.Trans)),
			nil,
		)
		return
	}

	tokenString, err := logic.Login(p)

	if err != nil {
		utils.ResponseWithCustomMsg(c, constants.CodeServerBusy, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(c, gin.H{
		"token": tokenString,
	})

}
