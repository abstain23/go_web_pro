package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gin-project/models"
)

func CheckUserExist(username string) (bool, error) {

	sqlStr := `select count(user_id) from user where username = ?`

	var count int

	err := db.Get(&count, sqlStr, username)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func InsertUser(user *models.User) (err error) {
	sqlStr := `insert into user(user_id, username, password, email, gender) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, encryptPassword(user.Password), user.Email, user.Gender)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	return
}

func encryptPassword(pwd string) string {
	h := md5.New()
	h.Write([]byte("go-project")) // secret
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}
