package models

type User struct {
	UserID     int64  `db:"user_id"`
	Username   string `json:"username" binding:"required,min=2" db:"username"`
	Password   string `json:"password" binding:"required" db:"password"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password" db:"password"`
	Email      string `json:"email" binding:"required" db:"email"`
	Gender     string `json:"gender" binding:"required" db:"gender"`
}
