package logic

import (
	"fmt"
	"gin-project/dao/redis"
	"gin-project/models"

	"go.uber.org/zap"
)

/*
	投票的几种情况：
		1. 之前没投过票，现在赞成票
		2. 之前投反对票，现在改投赞成票
		...

	投票的限制：
	每个帖子自发表之日起一个星期之内允许投票，超过不允许
		到期之后将redis中保存的赞成票数及反对票数放到mysql中
		到期之后删除那个保存的 （KetPostVotedZSetPrefix）

*/

// 投票功能
func VoteForPost(userID int64, p *models.ParamsVoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.String("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(fmt.Sprintf("%d", userID), p.PostID, float64(p.Direction))
}
