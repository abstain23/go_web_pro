package constants

type ResCode int64

const (
	CodeSuccess      ResCode = 0 + iota
	CodeInvalidParam ResCode = 1000 + iota
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

const UNKNOWN_ERROR = "未知错误"

var CodMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或者密码错误",
	CodeServerBusy:      "服务繁忙",
}
