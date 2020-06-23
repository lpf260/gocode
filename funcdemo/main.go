package main

import "fmt"

// 累加器
func AddUpper() func (int) int {
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
}

func main(){
	f := AddUpper()
	fmt.Printf("累加器求和：%d", f(20))
	fmt.Printf("累加器求和：%d", f(20))
	fmt.Printf("累加器求和：%d", f(20))

	fmt.Printf("累加器求和：%d", AddUpper()(20))
	fmt.Printf("累加器求和：%d", AddUpper()(20))
	fmt.Printf("累加器求和：%d", AddUpper()(20))
}