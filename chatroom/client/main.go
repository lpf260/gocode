package main

import (
	"fmt"
	"os"
)

// 
var userId int
var userPwd string

func main() {
	
	// 接受用户的选择
	var key int

	// 判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("---------------欢迎登录多人聊天室-----------------")
		fmt.Println("\t\t\t 1登录聊天室")
		fmt.Println("\t\t\t 2注册用户")
		fmt.Println("\t\t\t 3退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1 :
			fmt.Println("登录聊天室")
			loop = false
		case 2 :
			fmt.Println("注册用户")
			loop = false
		case 3 :
			fmt.Println("退出系统")
			os.Exit()
		default :
			fmt.Println("输入有误，请重新输入")
		}
	}

	// 根据用户的输入，显示新的提示信息
	if key == 1 {
		fmt.Println("请输入用户ID")
		fmt.Scanf("%d\n", &userId)
		
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)

		// 登录
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		}else{
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑...")
	}
}