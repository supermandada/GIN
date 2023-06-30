package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUpUser) (err error) {
	// 业务逻辑处理，判断用户是否存在
	if err := mysql.CheckUsernameExist(p.Username); err != nil {
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

func LogIn(p *models.ParamSignInUser) (err error) {
	return err
}
