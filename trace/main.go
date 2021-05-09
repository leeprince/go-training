package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

/*
go tool trace 跟踪执行器:
可以揭示Go程序运行中的所有的运行时事件, 是Go生态系统中用于诊断性能问题时（如延迟，并行化和竞争异常）最有用的工具之一。
*/
/** 使用
1. 添加 recordTrace 方法，该方法通用
2. 运行程序生成 *.out 文件
3. 执行 go tool trace {生成的*.out文件}
	此命令会自动创建监听端口并跳转到浏览器打开此监听端口。
	注意：部分视图页面依赖：Graphviz， 需安装才可查看
 */
var wg sync.WaitGroup

func main() {
	/*
	注意：将runtime/trace 的方法放在 recordTrace() 方法中并未能记录go 程序运行的事件！！因为包含文件的关闭：defer f.Close() 及trace的停止：defer trace.Stop()
	
	正确：将 recordTrace() 方法的内容全部放到main的方法体内
	 */
	// recordTrace()
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	
	// 测试协程一
	/*ch := make(chan int)
	done := make(chan bool)
	go send(1, 4, ch)
	go recv(ch, done)
	<-done*/
	
	// 测试协程二
	ReadMulFile()
}

func send(start, end int, ch chan<- int) {
	for i := start; i <= end; i++ {
		ch <- i
	}
	close(ch)
}

func recv(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Println("num : ", num)
	}
	done <- true
}


func ReadMulFile()  {
	ch := make(chan string) // 存储数据
	done := make(chan []string) // 启动阻塞, 达到自动切换协程
	
	wg.Add(3)
	go file(ch)
	go mysql(ch)
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

/* 通用方法
生成 *.out 文件
*/
func recordTrace() {
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
}
