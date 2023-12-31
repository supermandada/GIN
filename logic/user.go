package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUpUser) (err error) {
	// 业务逻辑处理，判断用户是否存在
	if ok, err := mysql.CheckUsernameExist(p.Username); ok {
		return err
	}
	// 生成UID
	var userID int64
	userID = snowflake.GenID()
	// 创建用户实例存入数据中
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 数据存入数据库
	return mysql.InsertUser(user)
}

func LogIn(p *models.ParamLogInUser) (token string, err error) {
	// 业务逻辑处理，判断用户是否存在
	if _, err := mysql.CheckUsernameExist(p.Username); err != mysql.ErrorUserExist {
		return "", err
	}
	// 创建用户实例存入数据中
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 判断密码是否正确
	if err := mysql.CheckPassword(user); err != nil {
		return "", err
	}
	// 生成jwt的token
	return jwt.GenToken(user.UserID, user.Username)
}
