package logic

import (
	"gin-project/dao/mysql"
	"gin-project/models"
	"gin-project/pkg/snowflake"
)

func CreatePost(post *models.Post) (err error) {

	post.ID = snowflake.GenID()
	return mysql.CreatePost(post)
}
