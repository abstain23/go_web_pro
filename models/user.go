package models

type User struct {
	UserID int64 `db:"user_id"`
	ParamsRegister
}
