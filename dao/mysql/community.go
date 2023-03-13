package mysql

import (
	"database/sql"
	"errors"
	"gin-project/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	if err = db.Select(&data, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			return make([]*models.Community, 0), nil
		}
	}

	return
}

func GetCommunityDetailById(id int64) (data *models.CommunityDetail, err error) {
	data = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time, update_time from community where community_id=?`
	if err = db.Get(data, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("无效的CommunityId")
		}
	}
	return
}
