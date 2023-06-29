package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/controllers"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/router"
	"web_app/settings"

	"go.uber.org/zap"
)

func main() {
	// go的web开发脚手架
	var filepath string
	flag.StringVar(&filepath, "file", "./conf/config.yaml", "项目配置文件")
	flag.Parse()

	//if len(os.Args) < 2 {
	//	fmt.Println("还没传参数。。。")
	//	return
	//}
	// 1、加载配置
	if err := settings.Init(filepath); err != nil {
		fmt.Println("settings init failed err:", err)
		zap.L().Error("settings init failed err:", zap.Error(err))
		return
	}
	// 2、初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("logger init failed err:", err)
		zap.L().Error("logger init failed err:", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 3、初始化mysql连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		zap.L().Error("mysql init failed err:", zap.Error(err))
		return
	}
	defer mysql.Close()
	// 4、初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("redis init failed err:", zap.Error(err))
		return
	}
	defer redis.Close()

	// 初始化雪花算法ID生产器
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		zap.L().Error("snowflake init failed err:", zap.Error(err))
		return
	}

	// 初始化错误翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Error("Trans init failed err:", zap.Error(err))
		return
	}
	// 5、注册路由
	r := router.Setup()
	// 6、启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
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
