package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}){
	// 通过反射获取传入变量的type，kind，值
	//1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 下面我们将rVal转成interface{}
	iV := rVal.Interface()
	// 需要断言才不会报错
	fmt.Printf("iV = %v iV=%T name=%v\n", iV, iV, iV.Name) // 编译阶段不能确定iV的类型，但是运行的时候是知道，反射是运行时反射
}

type Student struct {
	Name string
	Age int
}

func main(){
	// 演示反射操作

	// 1. 先定义一个int
	var num int = 100 // 通过反射拿到num的type和value
	reflectTest01(num)

	// 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age: 20,
	}

	reflectTest01(stu)
}