package routes

import (
	"gin-project/controllers"
	"gin-project/logger"
	"gin-project/middleware"
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

	// 登录
	r.POST("/login", controllers.LoginHandler)

	r.GET("/ping", middleware.JWTAuthMiddleware(), func(ctx *gin.Context) {
		username := ctx.GetString("username")
		ctx.JSON(http.StatusOK, gin.H{
			"data":     "【ping】【pong】~~",
			"username": username,
		})
	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "404 not found",
		})
	})
	return r
}
