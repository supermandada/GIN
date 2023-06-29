package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUpUser) {
	// 业务逻辑处理，判断用户是否存在
	mysql.QueryUserByUsername()
	// 生成UID
	snowflake.GenID()
	// 数据存入数据库
	mysql.InsertUser()
}
