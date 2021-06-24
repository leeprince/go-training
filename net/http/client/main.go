package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:8400/init")
	fmt.Println("server reply:", resp)
	
	data := make([]byte, 1024)
	index, _ := resp.Body.Read(data[:])
	fmt.Println("server reply body:", string(data[:index]))
}
