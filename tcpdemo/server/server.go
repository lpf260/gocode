package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	// 循环接受客户端发送的数据
	defer conn.Close() // 关闭链接

	for {
		buf := make([]byte, 1024)

		// 等待客户端通过conn发送信息
		// 如果客户端没有write[发送], 那么协程就会阻塞在这里

		fmt.Println("服务器在等待客户端发送信息" + conn.RemoteAddr().String())

		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			fmt.Println("服务器的Read err=", err)
			return

		}

		// 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main(){
	fmt.Println("服务器开始监听。。。")

	// 1. tcp表示使用网络协议是tcp
	
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return 
	}

	defer listen.Close() 

	for {
		// 等待客户端的链接
		fmt.Println("等待客户端来链接。。。")

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		}else{
			fmt.Printf("Accept() suc con=%v 客户端ip=%v \n", conn, conn.RemoteAddr())
		}

		// 这里准备一个协程，为客户端服务
		go process(conn)
	}
}