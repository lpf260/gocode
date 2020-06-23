package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string){
	// 判断users中是否有name
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	}else{
		// 没有这个用户
		users[name] = make(map[string]string,2)
		users[name]["pwd"] = "11111"
		users[name]["nickname"] = "昵称: " + name
	}
}

func main(){
	users := make(map[string]map[string]string, 10)
	modifyUser(users, "tom")
	modifyUser(users, "mary")

	fmt.Println(users)
}