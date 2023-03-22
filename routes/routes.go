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

	v1 := r.Group("/api/v1")
	v2 := r.Group("/api/v2")

	// 注册
	v1.POST("/register", controllers.RegisterHandler)

	// 登录
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middleware.JWTAuthMiddleware())

	{
		v1.GET("/community", controllers.CommunityListHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
		v1.POST("/community", controllers.CreatePostHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)

		v1.POST("/vote", controllers.PostVoteHandler)
	}

	v2.Use(middleware.JWTAuthMiddleware())

	{
		v2.POST("/posts", controllers.GetPostListHandlerV2)
	}

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
