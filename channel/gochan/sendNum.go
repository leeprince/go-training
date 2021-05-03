package gochan

import "fmt"

/*
有两个协程方法；
	发送方法可以循环发送指定范围的数字，
	接收可以接收并打印出send新的；
	并且程序必须等两个协程完成后再结束
*/
func SendNum() {
	ch := make(chan int)
	startNum := 1
	endNum := 20000
	
	// 方法一
	oneGoroutine(ch, startNum, endNum)
	
	// 方法二
	// twoGoroutine(ch, startNum, endNum)
	
	/*
	以上两种方法测试消耗时间基本一致
	 */
}

func oneGoroutine(ch chan int, startNum, endNum int) {
	go send1(ch, startNum, endNum)
	recv1(ch)
}

func twoGoroutine(ch chan int, startNum, endNum int) {
	done := make(chan bool)
	
	go send2(ch, startNum, endNum)
	go recv2(ch, done)
	<- done
}

func send2(ch chan<- int, start, end int) {
	for i := start; i <= end; i++ {
		ch <- i
		// fmt.Println("send Num:", i)
	}
	close(ch)
}

func recv2(ch <-chan int, done chan<- bool)  {
	for i := range ch {
		fmt.Println("recv Num:", i)
	}
	done <- true
}

func send1(ch chan<- int, start, end int) {
	for i := start; i <= end; i++ {
		ch <- i
		// fmt.Println("send Num:", i)
	}
	close(ch)
}

func recv1(ch <-chan int) {
	for {
		if n, bool := <-ch; bool {
			fmt.Println("recv Num:", n)
		} else {
			break
		}
	}
}
