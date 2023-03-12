package logic

import (
	"gin-project/dao/mysql"
	"gin-project/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	return mysql.GetCommunityList()
}
