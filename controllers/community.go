package controllers

import (
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 和社区相关的

func CommunityHandle(c *gin.Context) {
	// 获取community_id和community_name
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseErr(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
