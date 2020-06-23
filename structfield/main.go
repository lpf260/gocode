package main

import "fmt"

// 如果结构体的字段类型是：指针 slice 和map的零值都是nil，即还没有分配空间
// 如果需要使用这样的字段，需要先make，才能使用

type Person struct{
	Name string
	Age int
	Scores [5]float64
	ptr *int // 指针
	slice []int // 切片
	map1 map[string]string 
}
func main(){
	var p1 Person
	p1.slice = make([]int,10) // 使用切片一定要先make分配空间
	p1.slice[0] = 100

	fmt.Println(p1)
}