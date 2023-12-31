package models

// 定义获取参数的字段类型

// ParamSignUpUser 注册
type ParamSignUpUser struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogInUser 登录
type ParamLogInUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
