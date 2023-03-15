package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-project/models"
)

func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post
	(post_id, title, content, author_id, community_id)
	values (?,?,?,?,?);
	`
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityId)

	return
}

func GetPostById(id int64) (data *models.Post, err error) {
	sqlStr := `select post_id, author_id, community_id, status, title, content, create_time
	from post where post_id = ?`
	data = new(models.Post)

	if err = db.Get(data, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("无效的Post Id")
		}
	}
	fmt.Printf("err: %v\n", err)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {

	sqlStr := `select post_id, author_id, community_id, status, title, content, create_time
	from post limit ? offset ?`

	posts = make([]*models.Post, 0, 2)

	err = db.Select(&posts, sqlStr, size, (page-1)*size)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	return
}

func GetPostCount() (count int64) {
	sqlStr := `select count(*) from post`

	db.Get(&count, sqlStr)

	return

}
