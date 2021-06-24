package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start udp ...0.0.0.0:8300")
	// 1. 监听端口，传递 net.UDPAddr对象
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8300,
		Zone: "",
	})
	if err != nil {
		fmt.Println("conn err:", err)
	}
	defer conn.Close()
	
	/*
		以下三种方法收发信息；推荐使用 ReadFromUDPAndWriteToUDP 方法
			ReadFromUDPAndWriteToUDP() 收发正常
			ReadFromAndWriteTo() 收发正常
			ReadAndWrite() 收正常，发不正常
	*/
	// 读取到客户端消息：ReadFromUDP；发送给客户端呢消息：WriteToUDP
	ReadFromUDPAndWriteToUDP(conn)
	
	// 读取到客户端消息：ReadFrom；发送给客户端呢消息：WriteTo
	// ReadFromAndWriteTo(conn)
	
	// 读取到客户端消息：Read；发送给客户端呢消息：Write
	/*
		读取到客户端消息：Read 正常；
		发送给客户端呢消息：Write 不能正常发送消息给到客户端。具体原因待解释
	*/
	// ReadAndWrite(conn)
	/*
		以下三种方法收发信息；- end
	*/
}

func ReadFromUDPAndWriteToUDP(conn *net.UDPConn) {
	fmt.Println("new connection >")
	for {
		// 2. 接收信息
		data := make([]byte, 1024)
		index, udpAddr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			continue
		}
		fmt.Println("read client content: ", string(data[:index]))
		
		// 3. 发送消息
		conn.WriteToUDP([]byte("server send to your"), udpAddr)
	}
}

func ReadFromAndWriteTo(conn *net.UDPConn) {
	fmt.Println("new connection >")
	for {
		// 2. 接收信息
		data := make([]byte, 1024)
		index, udpAddr, err := conn.ReadFrom(data[:])
		if err != nil {
			continue
		}
		fmt.Println("read client content: ", string(data[:index]))
		
		// 3. 发送消息
		conn.WriteTo([]byte("server send to your"), udpAddr)
	}
}

func ReadAndWrite(conn *net.UDPConn) {
	fmt.Println("new connection >")
	for {
		// 2. 接收信息
		data := make([]byte, 1024)
		index, err := conn.Read(data[:])
		if err != nil {
			continue
		}
		fmt.Println("read client content: ", string(data[:index]))
		
		// 3. 发送消息
		conn.Write([]byte("server send to your"))
	}
}
