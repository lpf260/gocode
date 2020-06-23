package main

import (
	"fmt"
	"strings"
)

 

func makeSuffix(suffix string) func (string) string {
	return func (filename string) string {
		if strings.HasSuffix(filename,suffix) {
			return filename 
		}
		return filename + suffix
	}
}

func main(){
	fmt.Println("文件名", makeSuffix(".jpg")("aaaa.jpg"))
	fmt.Println("文件名", makeSuffix(".jpg")("bbb"))
}