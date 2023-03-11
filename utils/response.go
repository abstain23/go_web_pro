package utils

import (
	"gin-project/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	{
		code: 业务代码
		msg: ’‘，
		data:
	}
*/

type ResponseData struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data any) {
	ResponseWithCode(c, constants.CodeSuccess, data)
}

func ResponseError(c *gin.Context, code constants.ResCode) {
	ResponseWithCode(c, code, nil)
}

func ResponseWithCode(c *gin.Context, code constants.ResCode, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: int(code),
		Msg:  GetMsg(code),
		Data: data,
	})
}

func ResponseWithCustomMsg(c *gin.Context, code constants.ResCode, msg, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: int(code),
		Msg:  msg,
		Data: data,
	})
}

func GetMsg(code constants.ResCode) string {

	msg, ok := constants.CodMsgMap[code]

	if !ok {
		return constants.UNKNOWN_ERROR
	}
	return msg
}
