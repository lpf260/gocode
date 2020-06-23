package main

import "fmt"
func main(){
	// 二维数组的遍历
	var arr3 = [2][3]int{{1,2,3},{4,5,6}}

	for i := 0; i<len(arr3); i++{
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v \t", arr3[i][j])
		}

		fmt.Printf("\n")
	}

	// for-range来遍历二维数组
	for _, v := range arr3 {
		fmt.Println(v)
	}
}