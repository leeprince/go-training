package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var sendChOnly chan<- string
var recvChOnly <-chan string

/** 协程通信
在实际开发中协程之间的任务虽然是不冲突，但是相互之间最终获取的数据有时需要用到或者相互之间也需要通信；

为了解决这个问题go中提出一种特殊的类型：通道（channel）来解决这个问题
*/

/**
通信是一种同步形式：通过通道，两个协程在通信（协程会和）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。

甚至可以在通道两端互相阻塞对方，形成了叫做死锁的状态。Go 运行时会检查并 panic，停止程序。死锁几乎完全是由糟糕的设计导致的。

无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道。

*/

/** 关闭通道与检测
通道可以被显式的关闭，尽管它们和文件不同：不必每次都关闭。只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。

ch := make(chan string)
defer close(ch)

v, ok := <-ch // ok is true if v received val
*/

/** 通道方向
有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。

var sendChOnly chan<- int // 单向发送通道
var recvChOnly <-chan int // 单向获取通道
 */

func main() {
	// testCh()
	
	// testChDeedLock()
	
	chDirection()
}

func testCh() {
	ch := make(chan string)
	// ch := make(chan string, 2) // 设置缓冲通道
	
	wg.Add(1)
	go goSend(ch)
	
	wg.Add(1)
	go goRecv(ch)
	
	wg.Wait()
}

func testChDeedLock() {
	/** 发生错误实例
	不设置缓存通道可能会发生死锁，在没有协程等待的情况下会发生死锁
	*/
	/*ch := make(chan string)
	send(ch)
	recv(ch)*/
	
	/** 解决死锁. 方法2
	1. 设置缓冲通道
	2. 通过协程堵塞等待
	*/
	// 解决方法1：设置缓冲通道
	/*ch := make(chan string, 1)
	send(ch)
	recv(ch)*/
	
	// 解决方法2：通过协程堵塞等待
	/*ch := make(chan string)
	wg.Add(1)
	go goSend(ch)
	wg.Add(1)
	go goRecv(ch)
	wg.Wait()*/
	
	/** 注意
	1. 发送通道信息 与 接收通道信息有先发送再接收的顺序
	2. 就算使用协程时，要注意发送通道信息 与 接收通道信息的顺序问题
	*/
	// 解决方法正确：都使用协程，顺序无关，不会发生死锁。
	/*ch := make(chan string)
	wg.Add(1)
	go goRecv(ch)
	wg.Add(1)
	go goSend(ch)
	wg.Wait()*/
	//  解决方法正确：接收通道使用协程，不会发生死锁。
	/*ch := make(chan string)
	go goRecv(ch)
	wg.Add(1)
	send(ch)
	wg.Wait() // wg.Wait() 不能放在send(ch) 前面 */
	//  解决方法错误：接收通道未使用协程，会发生死锁。就算发送通道消息使用协程还是发生死锁
	/*ch := make(chan string)
	recv(ch)
	wg.Add(1)
	go goSend(ch)
	wg.Wait()*/
	//  解决方法错误: 就算有缓冲通道，但是顺序错误
	/*ch := make(chan string, 1)
	recv(ch)
	send(ch)*/
	//  解决方法错误：send() 发送2次通道信息，但是未存在2次接收通道，发生死锁
	/*ch := make(chan string)
	wg.Add(1)
	go goRecv(ch)
	send(ch)
	send(ch)
	wg.Wait() // wg.Wait() 不能放在send(ch) 前面*/
	// 解决方法错误：send() 发送1通道信息，但是存在2次接收通道，发生死锁
	/*ch := make(chan string)
	wg.Add(1)
	go goRecv(ch)
	send(ch)
	recv(ch)
	wg.Wait() // wg.Wait() 不能放在send(ch) 前面
	*/
	/** 关闭通道
	针对上面问题可以提前关闭通道
	 */
	ch := make(chan string)
	wg.Add(1)
	go goRecv(ch)
	send(ch)
	close(ch)
	v, bool := <- ch
	fmt.Println("v:", v, "-bool:", bool)
	wg.Wait() // wg.Wait() 不能放在send(ch) 前面
	
	
}

func chDirection()  {
	/*ch := make(chan string, 1)
	sendChOnly = ch
	sendChOnly <- "s001"
	recvChOnly = ch
	fmt.Println("recvChOnly:", <-recvChOnly)*/
	
	ch := make(chan string, 1)
	go sendfChOnly(ch)
	go recvfChOnly(ch)
	time.Sleep(time.Second)
}

func sendfChOnly(ch chan<- string)  {
	ch <- "s001"
}

func recvfChOnly(ch <-chan string)  {
	fmt.Println("recvChOnly:", <-ch)
}

func send(ch chan<- string) {
	s1 := "p100001"
	ch <- s1
	fmt.Println("通道发送值<-：", s1)
}

func recv(ch <-chan string) {
	fmt.Println("<-通道接收值：", <-ch)
}

func goSend(ch chan<- string) {
	defer wg.Done()
	
	s1 := "p00001"
	ch <- s1
	fmt.Println("通道发送值<-：", s1)
}

func goRecv(ch <-chan string) {
	defer wg.Done()
	fmt.Println("<-通道接收值：", <-ch)
}
