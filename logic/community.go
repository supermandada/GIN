package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]models.Community, error) {
	// 从数据库中查询并返回communityList
	return mysql.GetCommunityList()
}
