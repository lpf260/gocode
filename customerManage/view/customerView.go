package main

import (
	"fmt"
	"customerManage/service"
	"customerManage/model"
)

type customerView struct {
	// 
	key string // 接受用户输入
	
	loop bool    // 
	
	customerService *service.CustomerService // 指针类型

}

// 显示所有的客户信息
func (this *customerView) list() {
	// 首先，获取到当前所有的客户资料
	customers := this.customerService.List()

	// 显示客户列表
	fmt.Println("-----------------------客户列表------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-----------------------客户列表完成--------------------\n\n")
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for this.loop {
		fmt.Println("---------------客户信息管理软件--------------")
		fmt.Println("               1 添 加 客 户")
		fmt.Println("               2 修 改 客 户")
		fmt.Println("               3 删 除 客 户")
		fmt.Println("               4 客 户 列 表")
		fmt.Println("               5 退       出")
		fmt.Print("请选择(1-5)：")
		
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添 加 客 户")
			break
		case "2":
			fmt.Println("修 改 客 户")
			break
		case "3":
			fmt.Println("删 除 客 户")
			break
		case "4":
			this.list()
			break
		case "5":
			this.loop = false
			break
		default:
			fmt.Println("输入有误，请重新输入...")
		}
		
	}
}

// 得到用户的输入，信息构建新的客户
func (this *customerView) add() {
	fmt.Println("----------------------------添加客户-------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""	
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := ""	
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""	
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""	
	fmt.Scanln(&email)

	// 创建一个新的Customer实例

}

func main(){
	// 实例化结构体
	var customerView = customerView{
		key: "",
		loop: true,
	}

	// 
	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}

