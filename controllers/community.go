package controllers

import (
	"strconv"
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

func CommunityDetailHandle(c *gin.Context) {
	// 获取前端传入的id
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		zap.L().Error("community detail get a invalid param for search", zap.Error(err))
		ResponseErr(c, CodeInvalidParam)
		return
	}
	// 获取community_id和community_name
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseErr(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
