package controllers

import (
	"fmt"
	"gin-project/constants"
	"gin-project/logic"
	"gin-project/models"
	"gin-project/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func GetPostListHandler(c *gin.Context) {

	var (
		page, size int64
	)

	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	page, err := strconv.ParseInt(pageStr, 10, 0)
	if err != nil {
		page = 1
	}

	size, err = strconv.ParseInt(sizeStr, 10, 0)

	if err != nil {
		size = 10
	}

	data, total, err := logic.GetPostList(page, size)

	if err != nil {
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, gin.H{
		"list":  data,
		"total": total,
	})
}

func GetPostListHandlerV2(c *gin.Context) {
	// var (
	// 	page, size int64
	// )

	// pageStr := c.Query("page")
	// sizeStr := c.Query("size")
	// orderBy := c.Query("order")

	// page, err := strconv.ParseInt(pageStr, 10, 0)
	// if err != nil {
	// 	page = 1
	// }

	// size, err = strconv.ParseInt(sizeStr, 10, 0)

	// if err != nil {
	// 	size = 10
	// }

	p := &models.ParamsPostList{
		Page:  1,
		Size:  10,
		Order: constants.OrderByTime,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandlerV2 with invalid params", zap.Error(err))
		utils.ResponseError(c, constants.CodeInvalidParam)
		return
	}

	data, total, err := logic.GetPostListV2(p)

	if err != nil {
		utils.ResponseError(c, constants.CodeServerBusy)
		return
	}

	utils.ResponseSuccess(c, gin.H{
		"list":  data,
		"total": total,
	})
}
