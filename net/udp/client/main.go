package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接服务端
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8300,
		Zone: "",
	})
	defer conn.Close()
	// 发送消息
	conn.Write([]byte("client send to your"))
	// 接收消息
	data := make([]byte, 1024)
	index, udpArr, err := conn.ReadFromUDP(data[:])
	if err != nil {
		fmt.Println("read udp error: ", err)
	}
	fmt.Println("server reply:", string(data[:index]), udpArr)
}
