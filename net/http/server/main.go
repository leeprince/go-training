package main

import (
	"fmt"
	"net/http"
)

const (
	httpHost string = "0.0.0.0:8400"
)

func main() {
	fmt.Println("start http...", httpHost)
	// 路由地址
	http.HandleFunc("/init", initController)
	// http.ListenAndServeTLS(...) // https
	// 监听服务
	http.ListenAndServe(httpHost, nil) // http
}

func initController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new connection >")
	// TODO: [处理逻辑] - prince_todo 2021/5/15 12:47
	fmt.Println("request remote addr", r.RemoteAddr)
	fmt.Println("request url", r.URL)
	fmt.Println("request mothod", r.Method)
	
	w.Write([]byte("server send to your"))
}
