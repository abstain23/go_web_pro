package controllers

import (
	"fmt"
	"gin-project/constants"
	"gin-project/logic"
	"gin-project/models"
	"gin-project/utils"

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

	p.AuthorID = userID

	err = logic.CreatePost(p)

	if err != nil {
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, nil)
}
