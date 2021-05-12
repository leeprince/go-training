package problem

import (
	"fmt"
	"github.com/leeprince/utils"
	"sync"
)

var wg sync.WaitGroup
func ConcurrentProbem() {
	path := "./"
	file := "f.txt"
	data := "prince"
	wg.Add(6)
	for i := 0; i <= 5; i++  {
		/* 关闭协程：不会出现竞态问题
			运行结果：
				WrtiteFile successfuly
				exist data: prince
				exist data: prince
				exist data: prince
				exist data: prince
				exist data: prince
		 */
		// oper(path, file, data)
		
		/* 打开协程：出现竞态问题。如果没能复现竞态问题，请先删除 f.txt 文件 再次运行
			运行结果：
				WrtiteFile successfuly
				WrtiteFile successfuly
				WrtiteFile successfuly
				exist data: princeprinceprince
				exist data: princeprinceprince
				exist data: princeprinceprince
		 */
		go oper(path, file, data)
	}
	wg.Wait()
}

func oper(path, file, data string) {
	defer wg.Done()
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
}
