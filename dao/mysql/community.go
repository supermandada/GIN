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

func GetCommunityDetailById(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	strSql := "select `community_id`,`community_name`,`introduction`,`create_time` from `community` where id=? "
	if err = db.Get(communityDetail, strSql, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("db get communityDetail have invalidID", zap.Error(err))
			err = ErrorInvalidID
		}
	}
	return
}
