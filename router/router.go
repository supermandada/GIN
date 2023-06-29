package router

import (
	"net/http"
	"web_app/controllers"
	"web_app/logger"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 用户注册
	r.POST("/signup", controllers.SignUpHandle)

	// 注册路由
	r.GET("/", func(c *gin.Context) {
		//time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "ok")
	})
	r.GET("/version", func(c *gin.Context) {
		//time.Sleep(time.Second * 5)
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
