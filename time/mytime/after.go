package mytime

import (
	"fmt"
	"time"
)

/*
实现 time.After() 方法
 */
func After() {
	done := make(chan bool)
	go TimeAfter(done)
	fmt.Println("自定义超时方法。时间到", <-done)
}

func TimeAfter(done chan bool)  {
	time.Sleep(time.Second * 1)
	done <- true
}


