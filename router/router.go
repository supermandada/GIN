package router

import (
	"fmt"
	"net/http"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middlewares"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 用户注册
	r.POST("/signup", controllers.SignUpHandle)

	// 用户登录
	r.POST("/login", controllers.LogInHandle)

	// 注册路由
	r.GET("/version", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//time.Sleep(time.Second * 5)
		c.String(http.StatusOK, settings.Conf.Version)
		uid, _ := controllers.GetCurrentUser(c)
		fmt.Println(uid)
	})
	// 没有匹配到路由的情况
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
