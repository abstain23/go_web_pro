package logic

import (
	"fmt"
	"gin-project/dao/mysql"
	"gin-project/models"
	"gin-project/pkg/snowflake"
)

func CreatePost(post *models.Post) (err error) {

	post.ID = snowflake.GenID()
	return mysql.CreatePost(post)
}

func GetPostDetailById(id int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostById(id)

	if err != nil {
		return
	}

	fmt.Printf("post.AuthorID: %v\n", post.AuthorID)
	user, err := mysql.GetUserById(post.AuthorID)

	if err != nil {
		return
	}

	community, err := mysql.GetCommunityDetailById(post.CommunityId)

	if err != nil {

		return
	}

	data = new(models.ApiPostDetail)
	data.Post = post
	data.Community = &models.Community{
		Name: community.Name,
		ID:   community.ID,
	}
	data.AuthorName = user.Username
	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, total int64, err error) {

	posts, err := mysql.GetPostList(page, size)

	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		postDetail, err := GetPostDetailById(post.ID)
		if err != nil {
			return data, total, err
		}
		data = append(data, postDetail)
	}

	total = mysql.GetPostCount()

	if err != nil {
		return data, total, err
	}

	return
}