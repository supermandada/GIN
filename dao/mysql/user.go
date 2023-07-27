package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"web_app/models"
)

const secret = "m13750890761@gmail.com"

func CheckUsernameExist(username string) (bool, error) {
	var count int
	sqlStr := "select count(`user_id`) from `user` where `username`=?"
	if err := db.Get(&count, sqlStr, username); err != nil {
		//fmt.Println(err)
		return true, err
	}
	if count > 0 {
		return true, ErrorUserExist
	}
	return false, ErrorUserNotExist
}

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 存入数据库
	sqlStr := "insert into `user` (`user_id`,`username`,`password`) values (?,?,?)"
	if _, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password); err != nil {
		return err
	}
	return

}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func CheckPassword(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select `user_id`,`username`,`password` from user where `username`=?"
	if err := db.Get(user, sqlStr, user.Username); err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
