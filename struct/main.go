package main

import "fmt"

// 定义一个Cat结构体
type Cat struct {
	Name string
	Age int
	Color string
}


func main(){
	// 创建一个Cat变量
	var cat1 Cat
	fmt.Println(cat1) // { 0 } 因为string类型打印出来是空，所以不显示，把上面都换成int，会看到{0 0 0}
	// 或者去掉string类型的会打印{0} 和上面对比发现左右是有空格的

	cat1.Name = "小黄"
	cat1.Age = 2
	cat1.Color = "green"
	fmt.Println(cat1) 
	fmt.Printf("cat1的地址%p  \n ",&cat1) 
	fmt.Printf("cat1.Name的地址%p  \n ",&cat1.Name) 
	fmt.Printf("cat1.Age%p    \n ",&cat1.Age) 
	fmt.Printf("cat1.Color%p   \n ",&cat1.Color) 
}