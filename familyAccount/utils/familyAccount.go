package utils

import (
	"fmt"
)

type FamilyAccount struct {
	// 声明必要的字段
	
	// 生命一个变量，接受用户输入的选项
	key string

	// 控制循环的变量
	loop bool

	balance float64
	
	money float64
	
	note string

	details string

	flag bool
}

// 编写工厂模式 返回实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key: "",
		loop: true,
		balance: 10000,
		money: 0,
		note: "",
		details: "",
		flag: false,
	}
} 

// 显示主菜单
func (this *FamilyAccount) MainMenu() {
	for this.loop {
		// 显示主菜单
		fmt.Println("-------------------家庭记账软件--------------------")
		fmt.Println("                   1. 收支明细 ")
		fmt.Println("                   2. 登记收入 ")
		fmt.Println("                   3. 登记支出 ")
		fmt.Println("                   4. 退出")
		fmt.Println("                   请选择 1-4：")

		fmt.Scanln(&this.key)
		
		switch this.key {
			case "1":
				this.showDetails()
				break
			case "2":
				this.income()
				fmt.Println("this.flag===", this.flag)
				break
			case "3":
				this.pay()
				break
			case "4":
				this.exit()
				break
			default:
				fmt.Println("请输入正确的选项")
				break		
		}
	}

}

// 将显示明细封装
func (this *FamilyAccount) showDetails() {
	fmt.Println("-------------------当前收支明细记录--------------------")
	if this.flag {
		fmt.Println(this.details)

	}else{
		fmt.Println("请去写一笔记录吧")

	}
}

// 将登记收入写成一个方法 和*FamilyAccount
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明: ")
	fmt.Scanln(&this.note)
	// 将收入情况写入details
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if(this.money > this.balance){
		fmt.Println("余额不足")

	}
	this.balance -= this.money
	fmt.Println("本次支出说明: ")
	fmt.Scanln(&this.note)
	// 将支出情况写入details
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 退出系统
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出系统吗？请输入y/n")

	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" {
		this.loop = false
	}
}