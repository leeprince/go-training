package main

import (
	"lock/resolve"
	"os"
	"runtime/trace"
)

/*
Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）.

在程序中，同一个变量多个goroutine去修改的时候，肯定是不允许同时修改的，同时修改肯定会出问题.
所以当一个goroutine在修改之前需要加锁，修改结束在解锁，这样别的goroutine就可以去修改了。
 */
func main() {
	/* trace 工具分析 */
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
	/* trace 工具分析 -end */

	// 存在竞态问题，也是并发问题，也称线程安全问题
	// problem.ConcurrentProbem()

	// 使用互斥锁解决竞态问题
	// resolve.Mutex()
	
	// 使用读写锁解决竞态问题
	resolve.RWMutex()
	
	
}
