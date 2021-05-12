package main

import (
	"fmt"
	"github.com/leeprince/utils"
)

func main() {
	fmt.Print("请输入任意字符：")
	input := utils.CRead()
	fmt.Printf("读取到的输入值：%s \n", input)
	
}
