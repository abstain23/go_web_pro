package redis

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值得多少分
)

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(pID int64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(ctx, KeyPostTimeZSet, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: fmt.Sprintf("%d", pID),
	})
	pipeline.ZAdd(ctx, KeyPostScoreZSet, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: fmt.Sprintf("%d", pID),
	})

	_, err := pipeline.Exec(ctx)

	if err != nil {
		zap.L().Error("redis CreatePost err", zap.Error(err))
	}
	return err
}

func VoteForPost(userID, postID string, direction float64) error {
	// 1. 判断投票的限制

	postTime := rdb.ZScore(ctx, KeyPostTimeZSet, postID).Val()
	fmt.Printf("postTime: %v\n", postTime)
	fmt.Printf("(float64(time.Now().Unix()) - postTime): %v\n", (float64(time.Now().Unix()) - postTime))
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	// 2. 更新帖子的分数
	// 先查询用户之前给帖子的投票纪录
	keyPost := KetPostVotedZSetPrefix + postID
	ov := rdb.ZScore(ctx, keyPost, userID).Val()
	var op float64
	if direction > ov {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(ov - direction)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, KeyPostScoreZSet, op*diff*scorePerVote, postID)

	if direction == 0 {
		pipeline.ZRem(ctx, keyPost, postID)
	} else {
		pipeline.ZAdd(ctx, keyPost, &redis.Z{
			Score:  direction,
			Member: userID,
		})
	}

	_, err := pipeline.Exec(ctx)

	return err
}
