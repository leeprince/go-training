package main

import (
	"fmt"
	"go-training/calculating-reflect-factory/ctrlconsole"
)

func main()  {
	for {
		data := ctrlconsole.Cal{}
		data.LeftNum = ctrlconsole.InputOfNum(true)
		data.Operator = ctrlconsole.InputOperator()
		data.RigtNum = ctrlconsole.InputOfNum(false)

		/*
		获取反射对象，并通过反射对象去执行具体的方法
		 */
		data.GetReflectOfCaculete()

		fmt.Printf("【计算所得结果】%f %s %f = %f \n", data.LeftNum, data.Operator, data.RigtNum, data.Result)
	}
}