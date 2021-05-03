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
	ch := make(chan string)
	data := make([]string, 0)
	done := make(chan []string)
	
	wg.Add(1)
	go file(ch)
	wg.Add(1)
	go mysql(ch)
	wg.Add(1)
	go redis(ch)
	
	
	go readData(ch, data, done)
	
	
	wg.Wait()
	close(ch)
	
	fmt.Println("data: ", data)
	fmt.Println("done: ", <-done)
}

func readData(ch <-chan string, data []string, done chan<- []string)  {
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
