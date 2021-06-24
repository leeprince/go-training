package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

/*
tcp 粘包
tcp在发送数据的时候因为存在数据缓存的关系，
对于数据在发送的时候在 短时间内 如果连续发送很多小的数据的时候就会有可能一次性一起发送
还有就是对于大的数据就会分开连续发送多次
*/
/*
tcp 粘包处理方案：固定包头 + 包体 协议
*/
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
	msg := "我爱中国--"
	sdata := dataProtocol(msg)
	for i := 0; i < 10; i++ {
		conn.Write(sdata)
	}
	// 3. 接收服务端回复消息
	var data [1024]byte
	index, _ := bufio.NewReader(conn).Read(data[:])
	fmt.Println("server reply: ", string(data[:index]))
}

/* 数据格式处理
tcp 粘包处理方案：固定包头 + 包体 协议
*/
func dataProtocol(msg string) []byte {
	len := len(msg)
	
	/* 固定包头：2个字节
	固定包头2个字节用户存储数据的长度 =  1～256^2=65536 长度
	*/
	bagByteLen := int16(len)
	
	/*
		包体
	*/
	pkg := new(bytes.Buffer)
	binary.Write(pkg, binary.BigEndian, bagByteLen) // 大端发送字节流
	/* … 其实是go的一种语法糖。
	1. 用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
	2. 用法是slice可以被打散进行传递。
	*/
	data := append(pkg.Bytes(), []byte(msg)...)
	return data
}
