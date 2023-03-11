package routes

import (
	"gin-project/controllers"
	"gin-project/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册
	r.POST("/register", controllers.RegisterHandler)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "【pong】")
	})
	return r
}
