package main

import "fmt"

func fbn(n int) ([]uint64) {
	// 声明一个切片 切片大小n 
	/*
	切片可以使用内置函数 make 创建，函数签名为：

	func make([]T, len, cap) []T
	其中T代表被创建的切片元素的类型。函数 make 接受一个类型、一个长度和一个可选的容量参数。 调用 make 时，内部会分配一个数组，然后返回数组对应的切片。

	var s []byte
	s = make([]byte, 5, 5)
	// s == []byte{0, 0, 0, 0, 0}
	当容量参数被忽略时，它默认为指定的长度。下面是简洁的写法：

	s := make([]byte, 5)
	*/

	if n == 0 {
		return  []uint64{}
	}

	if n == 1 {		
		return []uint64{1}
	}

	fbnSlice := make([]uint64, n)

	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i -1] + fbnSlice[i - 2]
	}

	return fbnSlice
}

func main(){
	// 可以接收一个n int
	// 能够将斐波那契数列放到切片中
	// 数列形式 1 1 2 3 5 8
	// for循环来存放斐波那契数列 0 => 1 1=>1
	fbnSlice := fbn(20)
	fmt.Println(fbnSlice)
}