package main

import (
	"fmt"
	"calculating-factory/ctrlconsole"
)

func main()  {
	for {
		leftNum := ctrlconsole.InputOfNum(true)
		operator := ctrlconsole.InputOperator()
		rigtNum := ctrlconsole.InputOfNum(false)

		// 注册工厂模式
		result := ctrlconsole.NewOperationFactory(operator).Calculate(leftNum, rigtNum)

		fmt.Printf("【计算所得结果】%f %s %f = %f \n", leftNum, operator, rigtNum, result)
	}
}