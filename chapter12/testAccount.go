package main

import (
	"fmt"
)

func main() {

	// 生命一个变量，接受用户输入的选项
	key := ""



	// 控制循环的变量
	loop := true

	balance := 10000.0
	money := 0.0
	note := ""
	details := "收支\t账户金额\t收支\t说明\t"

	flag := false

	for loop {
		// 显示主菜单
		fmt.Println("-------------------家庭记账软件--------------------")
		fmt.Println("                   1. 收支明细 ")
		fmt.Println("                   2. 登记收入 ")
		fmt.Println("                   3. 登记支出 ")
		fmt.Println("                   4. 退出")
		fmt.Println("                   请选择 1-4：")

		fmt.Scanln(&key)
		
		switch key {
			case "1":
				fmt.Println("-------------------当前收支明细记录--------------------")
				if flag {
					fmt.Println(details)

				}else{
					fmt.Println("请去写一笔记录吧")

				}
				break
			case "2":
				fmt.Println("本次收入金额：")
				fmt.Scanln(&money)
				balance += money
				fmt.Println("本次收入说明: ")
				fmt.Scanln(&note)
				// 将收入情况写入details
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
				flag = true

				break
			case "3":
				fmt.Println("本次支出金额：")
				fmt.Scanln(&money)
				if(money > balance){
					fmt.Println("余额不足")
					break
				}
				balance -= money
				fmt.Println("本次支出说明: ")
				fmt.Scanln(&note)
				// 将支出情况写入details
				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				flag = true
				break
			case "4":
				fmt.Println("你确定要退出系统吗？请输入y/n")

				choice := ""
				fmt.Scanln(&choice)
				if choice == "y" {
					loop = false
					break
				}
				break
			default:
				fmt.Println("请输入正确的选项")
				break		
		}
	}

	if !loop {
		fmt.Println("你退出了家庭记账软件的使用")
	}

}