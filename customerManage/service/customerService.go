package service

import (
	"customerManage/model"
)

// 完成对Customer的操作，包括增删改查
type CustomerService struct{
	customers []model.Customer // 切片 装的是Customer结构体
	
	// 声明一个字段，表示当前切片含有多少个客户
	customerNum int
}

// 编写一个方法，返回*CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，默认一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 0

	customer := model.NewCustomer(	1, "张三", "男", 20, "112", "zs@sohu.com" )
	

	customerService.customers = append(customerService.customers, customer)

	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 增加客户
func (this *CustomerService) Add(customer Customer) bool {
	this.customers = append(this.customers, customer)
	return true
}