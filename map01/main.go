package main

import "fmt"
func main()  {
	var a map[string]string // map分配空间才能使用

	// 在使用map前，先make，make的作用就是给map分配数据空间

	a = make(map[string]string, 10)
	a["no1"] = "松江"
	a["no2"] = "松1111江"
	a["no3"] = "松222江"
	a["no4"] = "松13333江"
	fmt.Println(a)
}