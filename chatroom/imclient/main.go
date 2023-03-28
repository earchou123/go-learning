package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {
	// Login(1, "123")
	var key int
	var loop = true

	for loop {
		fmt.Println("-----------欢迎进入多人聊天系统-----------")
		fmt.Printf("%21s\n", "1 登录聊天室")
		fmt.Printf("%20s\n", "2 注册用户")
		fmt.Printf("%20s\n", "3 退出系统")
		fmt.Printf("%s\n", "请选择（1-3）")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录")
			loop = false
		case 2:
			fmt.Println("注册")
			loop = true
		case 3:
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}

	if key == 1 {
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPwd)

		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败", err)
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进入注册逻辑")
	}
}
