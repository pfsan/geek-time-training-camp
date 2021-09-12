package main

import (
	"errors"
	"fmt"
	"geek-time-training-camp/go/week02/dao"
)

func main() {
	// GetUserInfo 获取用户信息
	user, err := dao.GetUserInfo(10)

	// 判断错误与输出用户信息

	if errors.As(err, &dao.DataEmptyError) {
		fmt.Println(err)
	} else {
		fmt.Println("err", err)
	}
	
	fmt.Println("user", user)
}
