package main


func main()  {
	type myInt int // 给int取了别名，在go中myInt和int虽然都是int类型，但是go认为myInt和int为另两个类型
}