package main

import (
	"context"
	"fmt"
	"gin-project/dao/mysql"
	"gin-project/dao/redis"
	"gin-project/logger"
	"gin-project/pkg/snowflake"
	"gin-project/pkg/validator"
	"gin-project/routes"
	"gin-project/settings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed")
		return
	}
	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Println("init logger failed")
		return
	}

	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	// 3. 初始化MYSQL链接
	defer mysql.Close()
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Println("init db failed")
		return
	}

	// 4. 初始化redis链接
	defer redis.Close()
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("init redis failed")
		return
	}

	// 初始化雪花id
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachinedID); err != nil {
		fmt.Printf("init snowflake failed, err: %v", err)
		return
	}

	// 初始化翻译器
	if err := validator.InitTrans(settings.Conf.Lang); err != nil {
		fmt.Printf("init validator trans failed, err: %v", err)
		return
	}

	// 5. 注册路由
	r := routes.Setup(settings.Conf.Mode)
	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
