package controllers

import (
	"gin-project/constants"
	"gin-project/logic"
	"gin-project/utils"

	"github.com/gin-gonic/gin"
)

func CommunityListHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()

	if err != nil {
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, gin.H{
		"list": data,
	})

}
