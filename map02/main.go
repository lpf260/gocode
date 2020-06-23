package main

import (
	"fmt" 
)

func modify(map1 map[int]int){
	// map是引用类型，在函数里面修改，会影响map的值
	map1[10] = 900
}
func main(){
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2

	modify(map1)
	fmt.Println(map1)
}