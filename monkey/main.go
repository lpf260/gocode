package main

import "fmt"

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数有误")
	}

	if n == 10 {
		return 1 
	}else{
		return (peach(n + 1) +  1) * 2
	}
}

func main() {
	fmt.Printf("一共有%d个桃子", peach(1))
}