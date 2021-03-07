package ctrlconsole

import "fmt"

func InputOfNum(isReturn bool) float64 {
	if isReturn {
		fmt.Println("=================================================== \n")
	}

	fmt.Print("【请输入计算数字】")

	var i float64
	fmt.Scan(&i)

	return i
}
