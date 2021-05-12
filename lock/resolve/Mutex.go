package resolve

import (
	"fmt"
	"github.com/leeprince/utils"
	"sync"
)

var wg sync.WaitGroup

/*
Mutex是一个互斥锁，可以创建为其他结构体的字段；
零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。
 */
var mutex sync.Mutex

func Mutex() {
	path := "./"
	file := "f.txt"
	data := "prince"
	wg.Add(5)
	for i := 1; i <= 5; i++  {
		go oper(path, file, data)
	}
	wg.Wait()
}

func oper(path, file, data string) {
	defer wg.Done()
	
	mutex.Lock() // 加互斥锁
	
	rdata, _ := utils.ReadFileOfString(path+file)
	if rdata != "" {
		fmt.Println("exist data:", rdata)
		data = rdata
	} else {
		bool, err := utils.WrtiteFile(path, file, data)
		if bool == false || err != nil {
			fmt.Println("WrtiteFile err.", bool, err)
		}
		fmt.Println("WrtiteFile successfuly")
	}
	
	mutex.Unlock() // 解互斥锁
	
}
