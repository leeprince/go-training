package resolve

import (
	"fmt"
	"github.com/leeprince/utils"
	"sync"
)

/*
RWMutex是读写互斥锁。
该锁可以被同时多个读取者持有或唯一个写入者持有。
RWMutex可以创建为其他结构体的字段；零值为解锁状态。
RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。
 */
var rwmutex sync.RWMutex

func RWMutex()  {
	path := "./"
	file := "f.txt"
	data := "prince"
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		// go operWMutex(path, file, data) // 写锁解决竞态问题
		go operWRMutex(path, file, data) // 读写锁解决竞态问题
	}
	wg.Wait()
}

func operWMutex(path, file, data string) {
	defer wg.Done()
	
	rwmutex.Lock() // 加写锁
	
	rdata, _ := utils.ReadFileOfString(path+file)
	if rdata != "" {
		// fmt.Println("exist data:", rdata)
		data = rdata
	} else {
		bool, err := utils.WrtiteFile(path, file, data)
		if bool == false || err != nil {
			fmt.Println("WrtiteFile err.", bool, err)
		}
		// fmt.Println("WrtiteFile successfuly")
	}
	
	rwmutex.Unlock() // 解写锁
	
}

func operWRMutex(path, file, data string) {
	defer wg.Done()
	
	// TODO: [错误] - prince_add_todo
	rwmutex.RLock() // 加读锁
	rdata, _ := utils.ReadFileOfString(path+file)
	rwmutex.RUnlock() // 解读锁
	if rdata != "" {
		fmt.Println("exist data:", rdata)
		data = rdata
	} else {
		rwmutex.Lock() // 加写锁
		
		bool, err := utils.WrtiteFile(path, file, data)
		if bool == false || err != nil {
			fmt.Println("WrtiteFile err.", bool, err)
		}
		fmt.Println("WrtiteFile successfuly")
		
		rwmutex.Unlock() // 解写锁
	}
	
	/*rwmutex.Lock() // 加写锁
	rdata, _ := utils.ReadFileOfString(path+file)
	rwmutex.Unlock() // 解写锁
	if rdata != "" {
		fmt.Println("exist data:", rdata)
		data = rdata
	} else {
		rwmutex.RLock() // 加读锁
		
		bool, err := utils.WrtiteFile(path, file, data)
		if bool == false || err != nil {
			fmt.Println("WrtiteFile err.", bool, err)
		}
		fmt.Println("WrtiteFile successfuly")
		
		rwmutex.RUnlock() // 解读锁
	}*/
	
	
}
