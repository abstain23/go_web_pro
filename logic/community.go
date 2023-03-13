package logic

import (
	"gin-project/dao/mysql"
	"gin-project/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailById(id)
}
