package gochan

import (
	"fmt"
	"sync"
)

/*
利用协程通道获取从多个文件中读取的值
 */

var wg sync.WaitGroup

func ReadMulFile()  {
	ch := make(chan string) // 存储数据
	done := make(chan []string) // 启动阻塞, 达到自动切换协程
	
	wg.Add(1)
	go file(ch)
	wg.Add(1)
	go mysql(ch)
	wg.Add(1)
	go redis(ch)
	
	go readData(ch, done)
	
	wg.Wait()
	close(ch)
	
	fmt.Println("done: ", <-done)
}

func readData(ch <-chan string, done chan<- []string)  {
	data := make([]string, 0) // 临时存储数据
	for {
		if dt, bool := <-ch; bool {
			data = append(data, dt)
			fmt.Println("data for:", data)
		} else {
			break
		}
	}
	done <- data
}

func file(ch chan<- string)  {
	defer wg.Done()
	ch <- "file context."
}
func mysql(ch chan<- string)  {
	defer wg.Done()
	ch <- "mysql context."
}
func redis(ch chan<- string)  {
	defer wg.Done()
	ch <- "redis context."
}
