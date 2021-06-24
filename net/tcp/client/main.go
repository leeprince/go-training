package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 1. 创建与服务端的连接
	conn, err := net.Dial("tcp", "127.0.0.1:8200")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}
	defer conn.Close()
	// 2. 向服务端发送消息
	/*
	测试服务端接收字节数：英文字符占1个字节；中文占3个字节。典型的UTF-8编码
	 */
	// conn.Write([]byte("client send to your"))
	conn.Write([]byte("我爱中国"))
	// 3. 接收服务端回复消息
	var data [1024]byte
	index, err := bufio.NewReader(conn).Read(data[:])
	if err != nil {
		fmt.Println("reader conn err: ", err)
		return
	}
	fmt.Println("server reply: ", string(data[:index]))
}
