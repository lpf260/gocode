package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis" // redis包
)

func main() {
	conn, err := redis.Dial("tcp", "47.104.27.158:6379")

	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return 
	}

	defer conn.Close() // 关闭

	if _, err := conn.Do("AUTH", "zaq12wsx"); err != nil {
		return
	}

	_, err = conn.Do("HSet", "user01", "name","jary02")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age",18)
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 因为返回的r是interface{}
	// 因为name对应的值是string，因此我们需要转换
	r, err := redis.String(conn.Do("HGet", "user01", "name"))
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}