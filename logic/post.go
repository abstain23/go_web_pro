package logic

import (
	"fmt"
	"gin-project/dao/mysql"
	"gin-project/dao/redis"
	"gin-project/models"
	"gin-project/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {

	post.ID = snowflake.GenID()
	err = mysql.CreatePost(post)
	if err != nil {
		return err
	}

	err = redis.CreatePost(post.ID)
	return
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

func GetPostListV2(p *models.ParamsPostList) (data []*models.ApiPostDetail, total int64, err error) {

	postIds, err := redis.GetPostIDsInOrder(p)

	if err != nil {
		return
	}

	if len(postIds) == 0 {
		zap.L().Warn("redis GetPostIDsInOrder success ids == 0")
		return
	}

	posts, err := mysql.GetPostListByIDs(postIds)

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
