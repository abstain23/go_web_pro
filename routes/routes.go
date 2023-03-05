package routes

import (
	"gin-project/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "【pong】")
	})
	return r
}
