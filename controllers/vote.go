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

func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamsVoteData)

	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			utils.ResponseError(c, constants.CodeInvalidParam)
			return
		}
		utils.ResponseWithCustomMsg(
			c,
			constants.CodeInvalidParam,
			innerValidator.RemoveTopStruct(errs.Translate(innerValidator.Trans)),
			nil,
		)
		return
	}

	userID, err := getCurrentUser(c)

	if err != nil {
		utils.ResponseError(c, constants.CodeEmptyToken)
		return
	}

	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, nil)
}
