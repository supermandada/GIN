package controllers

import (
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
			return
		}

	}
	//fmt.Println(p)
	// 业务逻辑的处理
	if err := logic.SignUp(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("signup logic handle failed", zap.Error(err))
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// SignInHandle 登录路由逻辑处理
func SignInHandle(c *gin.Context) {
	// 接收和校验参数
	p := new(models.ParamSignInUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("signIn handle invalid param", zap.Error(err))
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
			return
		}

	}

	// 业务逻辑的处理
	if err := logic.SignIn(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("signup logic handle failed", zap.Error(err))
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
