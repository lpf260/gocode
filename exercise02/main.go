package main

import ("fmt")

func f(n int) int {
	if n == 1 {
		return 3
	}else {
		return 2 * f(n-1) + 1
	}
}

func main(){
	fmt.Println("输出", fmt.Println(f(10)))
}