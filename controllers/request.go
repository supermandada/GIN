package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserID = "UserID"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserID)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
