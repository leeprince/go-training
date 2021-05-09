package main

import (
	"fmt"
	"time"
)

func main() {
	TickerReq()
}

/*
定时发送请求
 */
func TickerReq()  {
	tk := time.Tick(1e9) // 纳秒; le9 纳秒= 1 秒
	for {
		<- tk
		curl()
	}
}

func curl()  {
	fmt.Println("curl 请求")
}
