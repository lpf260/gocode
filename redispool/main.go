package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis" // redis包
)

var pool *redis.Pool

// Go程序会自动调用init()和main()
func init() {
	pool = &redis.Pool{
		MaxIdle: 8,
		MaxActive: 0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error){
			return redis.Dial("tcp", "47.104.27.158:6379")
		},
	}
}
func main() {
	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("AUTH", "zaq12wsx"); err != nil {
		return
	}

	_, err := conn.Do("set", "name", "tangmu cat")
	if err != nil {
		fmt.Println("conn.Do err=", err)
	}

	// 取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do get err=", err)
	}

	fmt.Println("r=", r)
}