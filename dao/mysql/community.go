package mysql

import (
	"database/sql"
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
