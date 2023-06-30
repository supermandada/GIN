package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	code: 1000,  //程序中的错误码
	msg: xx,     //提示信息
	data: xx,    //返回的数据
}
*/

type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data"`
}

// ResponseErr 处理失败返回错误
func ResponseErr(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseWithMsg 处理失败返回自定义错误
func ResponseWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 响应成功
func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
