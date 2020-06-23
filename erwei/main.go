package main

import "fmt"

func main(){
	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Printf("arr2的是%v \n", arr2)
	fmt.Printf("arr2[0]的地址是%p \n", &arr2[0])
	fmt.Printf("arr2[1]的地址是%p \n", &arr2[1])
	fmt.Printf("arr2[0][0]的地址是%p \n", &arr2[0][0])
	fmt.Printf("arr2[0][1]的地址是%p \n", &arr2[0][1])
	fmt.Printf("arr2[0][2]的地址是%p \n", &arr2[0][2])
	fmt.Printf("arr2[1][0]的地址是%p \n", &arr2[1][0])
	fmt.Printf("arr2[1][1]的地址是%p \n", &arr2[1][1])
	fmt.Printf("arr2[1][2]的地址是%p \n", &arr2[1][2])
}