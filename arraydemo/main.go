package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// 随机生产5个数， rand.Intn() 函数
	// 当我们得到随机数后，放到一个数组 int 数组
	var intArr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)
	// 反转打印 交换的次数是 len / 2, 倒数第一个和第二个元素交换，
	temp := 0
	for i := 0; i < len(intArr3) / 2; i++ {
		temp = intArr3[len(intArr3) - 1 - i] // 倒数的某个元素intArr3[len(intArr3) - 1 - i]  
		intArr3[len(intArr3) - 1 - i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println(intArr3)

	var slice []string = []string{"tom", "jerry", "jack"}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}