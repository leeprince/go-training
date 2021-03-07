package main

import (
	"fmt"
	"go-training/calculating-normal/ctrlconsole"
)

type Calculate struct {
	LeftNum float64
	Operator string
	RigtNum float64
	Result float64
}

func main() {
	for  {
		data := Calculate{}
		data.LeftNum = ctrlconsole.InputOfNum(true)
		data.Operator = ctrlconsole.InputOperator()
		data.RigtNum = ctrlconsole.InputOfNum(false)
		
		data.Result, _ = ctrlconsole.Calculating(data.LeftNum, data.RigtNum, data.Operator)
		
		fmt.Printf("【计算所得结果】%f %s %f = %f \n", data.LeftNum, data.Operator, data.RigtNum, data.Result)
	}
}
