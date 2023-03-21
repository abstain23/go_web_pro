package redis

const (
	KeyPostTimeZSet        = "project:post:time"  // 帖子以发帖时间
	KeyPostScoreZSet       = "project:post:score" // 帖子以投票的分数
	KetPostVotedZSetPrefix = "project:voted:"     // 记录用户及投票类型 参数是post id
)
