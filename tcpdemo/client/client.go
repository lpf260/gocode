package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if( err != nil ) {
		fmt.Println("链接错误")
	}else{
		fmt.Printf("conn=%v \n", conn)
	}

	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("readString err=", err)
	}

	// 再将line发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn.Write err=", err)
	}

	fmt.Printf("客户端发送了 %d 字节的数据，并退出", n)
}