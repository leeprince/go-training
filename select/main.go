package main

import (
	"fmt"
	"time"
)

/* select {case ... defualt ...}
1. 每个case都必须是一个通信
2. 所有channel表达式都会被求值
3. 所有被发送的表达式都会被求值
4. 如果任意某个通信可以进行，它就执行;其他被忽略。
5. 如果有多个case都可以运行，select会随机公平地选出一个执行。其他不会执行。否则执行default子句(如果有) 6. 如果没有default字句，select将阻塞，直到某个通信可以运行;Go不会重新对channel或值进行求值。
 */
/*
select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下：

 select {
    case <-chan1:
       // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
       // 如果成功向chan2写入数据，则进行该case处理语句
    default:
       // 如果上面都没有成功，则进入default处理流程
 }

*/
func main() {
	chMysql := make(chan string)
	chRedis := make(chan string)
	
	go mysql(chMysql)
	go redis(chRedis)
	
	/*
		同时从多个通道接收数据
	*/
	// 方法一
	/*chMysqlData := <-chMysql
	// mysql 逻辑处理
	fmt.Println("chMysqlData:", chMysqlData)
	chRedisData := <-chRedis
	// redis 逻辑处理
	fmt.Println("chRedisData:", chRedisData)*/
	
	// 方法二
	/* break label
	break label是终止它所归属的最内循环整个循环，
	然后跳到label处开始执行，它所对一个的循环之后的那部分code不会被执行
	*/
	// time.After() 会重新返回当前时间；故不能直接在循环中直接
	afterTime := time.After(time.Second * 2)
DATALOOP:
	for {
		select {
		case chMysqlData := <-chMysql:
			fmt.Println("chMysqlData:", chMysqlData)
			break DATALOOP
		case chRedisData := <-chRedis:
			fmt.Println("chRedisData:", chRedisData)
		case <-afterTime:
			fmt.Println("time out <<<<<<<<<<<")
			break DATALOOP
		default:
			// fmt.Println("select default")
			// break DATALOOP
		}
	}
	
}

func mysql(ch chan<- string) {
	time.Sleep(time.Second * 3)
	ch <- "mysql context."
}
func redis(ch chan<- string) {
	time.Sleep(time.Second * 1)
	ch <- "redis context."
}
