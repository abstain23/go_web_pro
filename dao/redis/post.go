package redis

import (
	"gin-project/constants"
	"gin-project/models"
)

func GetPostIDsInOrder(p *models.ParamsPostList) ([]string, error) {

	key := KeyPostTimeZSet

	if p.Order == constants.OrderByScore {
		key = KeyPostScoreZSet
	}

	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	return rdb.ZRevRange(ctx, key, start, end).Result()
}
