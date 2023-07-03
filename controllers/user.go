package controllers

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 路由处理 参数校验 请求转发

// SignUpHandle 注册路由逻辑处理
func SignUpHandle(c *gin.Context) {
	// 接收和校验参数
	p := new(models.ParamSignUpUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("signup handle invalid param", zap.Error(err))
		if !ok {
			ResponseErr(c, CodeInvalidParam)
			return
		}
		ResponseWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return

	}
	//fmt.Println(p)
	// 业务逻辑的处理
	if err := logic.SignUp(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("signup logic handle failed", zap.Error(err))
		if !ok {
			if errors.Is(err, mysql.ErrorUserExist) {
				ResponseErr(c, CodeUserExist)
				return
			}
			ResponseErr(c, CodeServerBusy)
			return
		}
		ResponseWithMsg(c, CodeServerBusy, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// LogInHandle 登录路由逻辑处理
func LogInHandle(c *gin.Context) {
	// 接收和校验参数
	p := new(models.ParamLogInUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("login handle invalid param", zap.Error(err))
		if !ok {
			ResponseErr(c, CodeInvalidParam)
			return
		}
		ResponseWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return

	}

	// 业务逻辑的处理
	token, err := logic.LogIn(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("login logic handle failed", zap.String("username", p.Username), zap.Error(err))
		if !ok {
			if errors.Is(err, mysql.ErrorInvalidPassword) {
				ResponseErr(c, CodeInvalidPassword)
				return
			}
			ResponseErr(c, CodeServerBusy)
			return
		}
		ResponseWithMsg(c, CodeServerBusy, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
}
