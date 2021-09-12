package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DataEmptyError = errors.New("data is empty")

type User struct {
	UserId   int64
	Username string
	Age      int64
}

// GetUserInfo 获取用户信息
func GetUserInfo(userId int64) (User, error) {
	var user User

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pfs?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.QueryRow(`
        SELECT user_id,username,age from user WHERE user_id = ?
    `, userId).Scan(
		&user.UserId, &user.Username, &user.Age,
	)
	switch {
	case err == sql.ErrNoRows:
		return user, DataEmptyError
	case err != nil:
		return user, fmt.Errorf("%w", err)
	}
	return user, nil
}
