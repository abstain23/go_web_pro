package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
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
		return err
	}
	return
}

func Login(user *models.User) (err error) {
	oldPasswd := user.Password

	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("用户不存在")
		}
		return err
	}

	if encryptPassword(oldPasswd) != user.Password {
		return errors.New("用户名或者密码不匹配")
	}

	return nil

}

func GetUserById(id int64) (user *models.User, err error) {
	sqlStr := `select username from user where user_id = ?`
	user = new(models.User)

	if err = db.Get(user, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("无效的User Id")
		}
	}
	return
}

func encryptPassword(pwd string) string {
	h := md5.New()
	h.Write([]byte("go-project")) // secret
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}
