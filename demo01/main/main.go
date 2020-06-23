package main

import (
	"fmt"
	// "demo01/model"
)
func main()  {
	var name string

	fmt.Println("请输入姓名")

	// fmt.Scanln(&name)

	// fmt.Println("name=", name)

	fmt.Scanf("%s", &name)

	fmt.Println("name=", name)

}