package main

// import相对路径从src开始计算
import (
	"familyAccount/utils"
)

func main() {
	utils.NewFamilyAccount().MainMenu()
}