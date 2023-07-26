package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []models.Community, err error) {
	strSql := "select `community_id`,`community_name` from `community`"
	if err := db.Select(&communityList, strSql); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("communityList is null in this db")
			err = nil
		}
	}
	return
}
