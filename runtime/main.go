package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
Go 中，协程被称为 goroutine，它非常轻量，一个 goroutine 只占几 KB，并且这几 KB 就足够 goroutine 运行完，这就能在有限的内存空间内支持大量 goroutine，支持了更多的并发。虽然一个 goroutine 的栈只占几 KB，但实际是可伸缩的，如果需要更多内容，runtime 会自动为 goroutine 分配。

Go 为了提供更容易使用的并发方法，使用了 goroutine 和 channel。goroutine 来自协程的概念，让一组可复用的函数运行在一组线程之上，即使有协程阻塞，该线程的其他协程也可以被 runtime 调度，转移到其他可运行的线程上。最关键的是，程序员看不到这些底层的细节，这就降低了编程的难度，提供了更容易的并发。
 */

/**
runtime
	Gosched() // 让当前线程让出 cpu 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行
	Goexit() // 退出当前 goroutine(但是defer语句会照常执行)
	GOMAXPROCS() // 设置最大的可同时使用的 CPU 核数
 */

/** 协程多核调度-GPM模型
G: 协程
P: 处理器（协程调度器）
M: 线程
 */
func main() {
	
	// runtime.GOMAXPROCS(1) // 方便演示runtime.Gosched()后的效果
	runtime.GOMAXPROCS(2)
	
	go A()
	go B()
	
	time.Sleep(time.Second * 1) // sleep 解决 goroutine main 主协程退出问题
}

func A()  {
	fmt.Println("A >")
	runtime.Gosched()
	fmt.Println("A > >")
}

func B()  {
	fmt.Println("B >")
	runtime.Goexit()
	fmt.Println("B > >")
}


