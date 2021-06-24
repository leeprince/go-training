package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
)

const websocketHost string = "127.0.0.1:8500"

func main() {
	fmt.Println("start websocket...", websocketHost)
	http.Handle("/", websocket.Handler(server))
	http.ListenAndServe(websocketHost, nil)
}

// 当有连接进来的时候底层会自动
func server(ws *websocket.Conn) {
	fmt.Println("new connection >")
	defer ws.Close()
	data := make([]byte, 1024)
	for {
		// 读取信息
		index, err := ws.Read(data[:])
		if err != nil {
			fmt.Println("err : ", err)
			break
		}
		fmt.Println("read client content ", string(data[:index]))
		ws.Write([]byte("server send to your"))
	}
}
