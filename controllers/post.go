package controllers

import (
	"fmt"
	"gin-project/constants"
	"gin-project/logic"
	"gin-project/models"
	"gin-project/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)

	if err := c.ShouldBindJSON(p); err != nil {
		fmt.Printf("err: %v\n", err)
		utils.ResponseError(c, constants.CodeInvalidParam)
		return
	}

	userID, err := getCurrentUser(c)

	if err != nil {
		utils.ResponseWithCustomMsg(c, constants.CodeInvalidToken, "不合法的用户", nil)
		return
	}

	fmt.Printf("userID: %v\n", userID)

	p.AuthorID = userID

	err = logic.CreatePost(p)

	if err != nil {
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseError(c, constants.CodeInvalidParam)
		return
	}

	data, err := logic.GetPostDetailById(id)

	if err != nil {
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, data)
}
