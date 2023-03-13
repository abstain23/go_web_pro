package controllers

import (
	"gin-project/constants"
	"gin-project/logic"
	"gin-project/utils"
	"strconv"

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

func CommunityDetailHandler(c *gin.Context) {
	cIDStr := c.Param("id")
	id, err := strconv.ParseInt(cIDStr, 10, 64)
	if err != nil {
		utils.ResponseError(c, constants.CodeInvalidCommunityID)
		return
	}

	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, data)
}
