package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
)

const (
	tcpHost    string = "0.0.0.0:8200"
	bagHeadLen int    = 2 // 固定包头字节数
)

/*
Tcp特点： 即传输控制协议/网间协议，是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议，因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。
	场景：短信、聊天

Udp特点：不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域。
	场景：直播、广播
*/
/*
tcp 粘包处理方案：固定包头 + 包体 协议
*/
var wg sync.WaitGroup
func main() {
	fmt.Println("start tcp ...", tcpHost)
	/*
		1. 监听端口 tcp:://0.0.0.0:8100
		返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络：
		"tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket"。参见Dial函数获取laddr的语法。
	*/
	listener, err := net.Listen("tcp", tcpHost)
	if err != nil {
		fmt.Println("listen err", err)
		return // return 表示程序结束
	}
	for {
		// 2. 接收客户端向服务端建立的连接
		/*
			Accept等待并返回下一个连接到该接口的连接
		*/
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err: ", err)
			return
		}
		// 3. 处理用户的连接信息
		wg.Add(1)
		go handler(conn)
	}
	wg.Wait()
}

/*
处理用户的连接信息
*/
func handler(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()
	fmt.Println("new connection >")
	reader := bufio.NewReader(conn)
	for { // 持续接收当前已建立连接的客户端输入
		/* 每次消息接收的字节数
		服务端接收字节数：英文字符占1个字节；中文占3个字节。典型的UTF-8编码
		*/
		// 读取客户端输入的内容
		len, err := unpack(reader)
		if err != nil {
			fmt.Println("unpack err: ", err)
			break // 结束持续接收客户端输入
		}
		data := make([]byte, len) // var data [6]byte
		index, err := reader.Read(data[:])
		/* 结束当前连接
		接收客户端端消息结束后，会返回一个错误，错误为结束符：EOF
		所以此处应判断错误并结束持续接收客户端输入
		否则会不断的接收到错误信息（结束符）
		*/
		if err != nil {
			fmt.Println("reader conn err: ", err)
			break // 结束持续接收客户端输入
		}
		fmt.Println("read content: ", string(data[:index]))
		// 向连接的客户端发送消息
		conn.Write([]byte("server send to your"))
	}
}

/* 数据解析
tcp 粘包处理方案：固定包头 + 包体 协议
*/
func unpack(reader *bufio.Reader) (int, error) {
	/* 包头解析
	reader.Peek 读取前 n 个字节数据
	*/
	/* 固定包头：2个字节
	固定包头2个字节用户存储数据的长度 =  1～256^2=65536 长度
	*/
	bagHeadLenData, _ := reader.Peek(bagHeadLen)
	bagHeadLenBuff := bytes.NewBuffer(bagHeadLenData)
	// 固定包头2个字节确定包体数据长度(2个字节表示)
	var bagBodyLen int16
	err := binary.Read(bagHeadLenBuff, binary.BigEndian, &bagBodyLen)
	if err != nil {
		fmt.Println("读取固定包头长度错误：", err)
		return 0, err
	}
	fmt.Println("bagBodyLen:", bagBodyLen)
	
	/*
		数据长度判断
	*/
	readDataLen := int16(reader.Buffered())
	dataLen := int16(bagHeadLen) + bagBodyLen
	if readDataLen < dataLen {
		return 0, errors.New("数据异常：数据长度不匹配")
	}
	// 返回数据长度
	return int(dataLen), nil
}

/*
int8: -127 ~ 127
uint8: 0 ~ 255
 */
func TestInt()  {
	var i int = 255
	var ii int = 128
	bi := int8(i)
	bii := int8(ii)
	ubi := uint8(i)
	ubii := uint8(ii)
	fmt.Println("i:", i, "bi:", bi, "bii:", bii, "ii:", ii, "ubi:", ubi, "ubii:", ubii)
}
