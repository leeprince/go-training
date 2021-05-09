package nested

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func nested() {
	startTime := time.Now().Unix()
	
	wg.Add(3)
	go redisWg()
	go mysqlWg()
	go fileWg()
	
	wg.Wait()
	
	endTime := time.Now().Unix()
	fmt.Printf("main 结束时间戳：%d; 总运行时间(单位：秒)：%d \n", endTime, endTime-startTime)
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
	
	// mysqlWgNestedNotGo()
	
	/* 协程嵌套
	协程嵌套保证所有协程运行完毕后，主协程（main）才可以退出的需注意：
		外部wg.Add(num) 需加上该协程做为总的协程个数==wg.Add(4),
		或者外部总的 wg.Add(3) 加上此处的 wg.Add(1)
	 */
	wg.Add(1)
	go mysqlWgNestedGo()
}

func mysqlWgNestedNotGo() {
	fmt.Println("mysqlWgNestedNotGo > start")
	time.Sleep(time.Second * 2)
	fmt.Println("mysqlWgNestedNotGo > end")
}

func mysqlWgNestedGo() {
	defer wg.Done() // 结束登记
	fmt.Println("mysqlWgNestedGo > start")
	time.Sleep(time.Second * 2)
	fmt.Println("mysqlWgNestedGo > end")
}

func fileWg() {
	defer wg.Done() // 结束登记
	fmt.Println("file > start")
	time.Sleep(time.Second * 3)
	fmt.Println("file > end")
}
