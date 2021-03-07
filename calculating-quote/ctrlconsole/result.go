// 包名的定义方式：父级目录（入口文件统一为main, 并且与入口文件同级的go文件包名同样为main包）
package ctrlconsole

import (
	"fmt"
)

// 变量名称、方法、函数 遵循：大写开头是允许外部访问；小写开头仅允许本包中方法
type Calculate struct {
	LeftNum float64
	Operator string
	RigtNum float64
	Result float64
}

func (data *Calculate) Calculating() { // 编号0001【推荐】
	method := data.Operator
	switch method {
	case "+":
		data.Result = addition(data.LeftNum, data.RigtNum)
	case "-":
		data.Result = subtraction(data.LeftNum, data.RigtNum)
	case "*":
		data.Result = multiplication(data.LeftNum, data.RigtNum)
	case "/":
		data.Result = division(data.LeftNum, data.RigtNum)
	default:
		fmt.Println("输入运算符异常")
	}
}

func Calculating(data *Calculate) { // 编号0002
	method := data.Operator
	switch method {
	case "+":
		data.Result = addition(data.LeftNum, data.RigtNum)
	case "-":
		data.Result = subtraction(data.LeftNum, data.RigtNum)
	case "*":
		data.Result = multiplication(data.LeftNum, data.RigtNum)
	case "/":
		data.Result = division(data.LeftNum, data.RigtNum)
	default:
		fmt.Println("输入运算符异常")
	}
}

func addition(l float64, r float64) float64  {
	return l + r
}
func subtraction(l float64, r float64) float64  {
	return l - r
}
func multiplication(l float64, r float64) float64  {
	return l * r
}
func division(l float64, r float64) float64  {
	return l / r
}