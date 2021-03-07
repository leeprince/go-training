package main

import (
	"fmt"
	"go-training/calculating-quote/ctrlconsole"
)
func main() {
	for  {
		data := ctrlconsole.Calculate{}
		data.LeftNum = ctrlconsole.InputOfNum(true)
		data.Operator = ctrlconsole.InputOperator()
		data.RigtNum = ctrlconsole.InputOfNum(false)

		/*
		同一个包中，方法与函数同名不会产生冲突：方法是结构体的函数
		 */
		data.Calculating() // 正常。编号0001 【推荐】
		// ctrlconsole.Calculating(&data) // 正常。编号0002
		fmt.Printf("【计算所得结果】%f %s %f = %f \n", data.LeftNum, data.Operator, data.RigtNum, data.Result)
	}
}
