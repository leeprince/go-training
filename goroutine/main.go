package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	startTime := time.Now().Unix()
	fmt.Println("main 开始时间戳：", startTime)
	
	redis()
	
	mysql()
	
	file()
	
	endTime := time.Now().Unix()
	fmt.Printf("main 结束时间戳：%d; 总运行时间(单位：秒)：%d \n", endTime, endTime-startTime)
	
	fmt.Println("⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎无协程")
	
	// 开启go协程方法：通过关键字 go 声明方法是一个协程。
	/*
		问题：协程逻辑未执行完，主程序退出
		原因：注意：main() 方法本身就是一个协程，当 main() 主协程结束时，其他协程不再执行
		解决：
			1. 外部添加协程中最大等待时间
			2. 登记协程：sync.WaitGroup
	*/
	startTimeGo := time.Now().Unix()
	fmt.Println("main 开启协程-开始时间戳：", startTimeGo)
	
	go redis()
	
	go mysql()
	
	go file()
	
	endTimeGo := time.Now().Unix()
	fmt.Printf("main 结束时间戳：%d; 总运行时间(单位：秒)：%d \n", endTimeGo, endTimeGo-startTimeGo)
	
	fmt.Println("⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎有协程：协程逻辑未执行完，主程序退出")

	startTimeGoSleep := time.Now().Unix()
	fmt.Println("main 开启协程-开始时间戳：", startTimeGo)
	
	go redis()
	
	go mysql()
	
	go file()
	
	time.Sleep(time.Second * 3)
	
	endTimeGoSleep := time.Now().Unix()
	fmt.Printf("main 结束时间戳：%d; 总运行时间(单位：秒)：%d \n", endTimeGoSleep, endTimeGoSleep-startTimeGoSleep)
	
	fmt.Println("⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎有协程：协程逻辑未执行完，主程序退出 ｜ sleep 解决")

	startTimeGoWaitGroup := time.Now().Unix()
	fmt.Println("main 开启协程-开始时间戳：", startTimeGo)
	// 登记协程数可一次性登记多个或所有协程；如一下有三个协程，可一次性写为：wg.Add(3)， 然后关闭各个wg.Add(1)。推荐单个协程登记协程
	// wg.Add(3) // 登记多个协程；
	wg.Add(1) // 登记协程；
	go redisWg()
	
	wg.Add(1) // 登记协程
	go mysqlWg()
	
	wg.Add(1) // 登记协程
	go fileWg()
	
	wg.Wait() // 等待登记的协程执行完再取消堵塞，执行后续的程序
	
	endTimeGoWaitGroup := time.Now().Unix()
	fmt.Printf("main 结束时间戳：%d; 总运行时间(单位：秒)：%d \n", endTimeGoWaitGroup, endTimeGoWaitGroup-startTimeGoWaitGroup)
	fmt.Println("⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎有协程：协程逻辑未执行完，主程序退出 ｜ 登记协程：sync.WaitGroup 解决")
}

func redis() {
	fmt.Println("redis > start")
	time.Sleep(time.Second * 1)
	fmt.Println("redis > end")
}

func mysql() {
	fmt.Println("mysql > start")
	time.Sleep(time.Second * 2)
	fmt.Println("mysql > end")
}
func file() {
	fmt.Println("file > start")
	time.Sleep(time.Second * 3)
	fmt.Println("file > end")
}

func redisWg() {
	defer wg.Done() // 结束登记
	fmt.Println("redis > start")
	time.Sleep(time.Second * 1)
	fmt.Println("redis > end")
}

func mysqlWg() {
	defer wg.Done() // 结束登记
	fmt.Println("mysql > start")
	time.Sleep(time.Second * 2)
	fmt.Println("mysql > end")
}
func fileWg() {
	defer wg.Done() // 结束登记
	fmt.Println("file > start")
	time.Sleep(time.Second * 3)
	fmt.Println("file > end")
}
