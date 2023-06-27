package router

import (
	"net/http"
	"time"
	"web_app/logger"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "ok")
	})
	r.GET("/version", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
