package main

import (
	"fmt"
)
type Point struct{
	x int
	y int
}

func main(){
	var a interface{}
	var point Point = Point{1, 2}
	a = point

	var b Point
	b = a.(Point)
	fmt.Printf("a type:%T\n b type:%T\n", a, b)
}