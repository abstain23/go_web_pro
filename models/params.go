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
