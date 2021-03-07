package ctrlconsole

import (
	"errors"
	"fmt"
)

func Calculating(l float64, r float64, p string) (float64, error) {
	switch p {
	case "+":
		return addition(l, r), nil
	case "-":
		return subtraction(l, r), nil
	case "*":
		return multiplication(l, r), nil
	case "/":
		return division(l, r), nil
	default:
		fmt.Println("输入运算符异常")
		return 0, errors.New("输入运算符异常")
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