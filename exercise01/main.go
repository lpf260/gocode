package main

import "fmt"

func fbn(n int) int{
	if (n == 1 || n == 2) {
		return 1
	}else{
		return fbn(n - 1) + fbn(n - 2)
	}

}

func main(){
	res := fbn(3)


	fmt.Println("[数列]", res)
	fmt.Println("[数列]", fbn(4))
	fmt.Println("[数列]", fbn(5))
	fmt.Println("[数列]", fbn(6))
}