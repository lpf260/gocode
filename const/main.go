package main

import (
	"fmt"
)

func main() {
	// 常量必须初始化
	const (
		a = iota
		b 
		c
		d 
	)

	fmt.Println(a, b,c,d)
}