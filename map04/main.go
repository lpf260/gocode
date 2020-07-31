package main

import "fmt"

func main() {
	// 相当于键是string类型，值是interface类型 直接map[sting]interface{}后面接大括号是为了直接赋值
	// 省去了make分配空间
	map01 := map[string]interface{}{
		"message": "pong",
		"status": 1,
	}

	fmt.Println(map01)
}

