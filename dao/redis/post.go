package redis

import (
	"gin-project/constants"
	"gin-project/models"

	"github.com/go-redis/redis/v8"
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

// 根据ids 查询投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	// data = make([]int64, 0, len(ids))

	// for _, id := range ids {
	// 	key := KetPostVotedZSetPrefix + id
	// 	v1 := rdb.ZCount(ctx, key, "1", "1").Val()
	// 	data = append(data, v1)
	// }

	pipeline := rdb.TxPipeline()
	for _, id := range ids {
		pipeline.ZCount(ctx, KetPostVotedZSetPrefix+id, "1", "1")
	}

	cmders, err := pipeline.Exec(ctx)

	if err != nil {
		return
	}

	data = make([]int64, 0, len(cmders))

	for _, cmd := range cmders {
		data = append(data, cmd.(*redis.IntCmd).Val())
	}

	return
}
