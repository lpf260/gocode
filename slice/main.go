package main

import "fmt"

func main(){
	var slice1 []int = []int{100,200,300}
	var slice2 []int = []int{400,500,600}

	slice3 := append(slice1,slice2...)
	fmt.Println(slice3)
}