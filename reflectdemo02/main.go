package main

import (
	"fmt"
	"reflect"
)


func reflectTest01(b interface{}){
	// 获取到reflect.Value
	rVal := reflect.ValueOf(b)

	// 看看rVal的kind
	fmt.Printf("rVal kind=%v \n", rVal.Kind())


	 rVal.Elem().SetInt(200) // rVal.Elem()有点类似于*ptr转指针的这种操作
}

// 通过反射，修改num int的值
// 修改student的值
func main() {
	var num int = 10
	reflectTest01(&num)

	fmt.Println("num=", num)


}

