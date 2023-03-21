package models

type ParamsRegister struct {
	Username   string `json:"username" binding:"required,min=2"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
}

type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamsVoteData struct {
	// UserID 从token拿
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"required,oneof=1 0 -1"` // 赞成票 1 反对票-1 取消票 0
}
