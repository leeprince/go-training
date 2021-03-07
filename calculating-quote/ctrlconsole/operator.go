package ctrlconsole

import "fmt"

type noRelOperator struct {
	no int
	char string
}

var countryCapitalMap = map[int]string{
	1: "+",
	2: "-",
	3: "*",
	4: "/"}

func InputOperator() string {
	fmt.Println("【请选择要操作符号对应的序号】")
	fmt.Println("> 1 （+）")
	fmt.Println("> 2 （-）")
	fmt.Println("> 3 （*）")
	fmt.Println("> 4 （/）")

	rel := noRelOperator{}
	fmt.Scan(&rel.no)

	value, ok := countryCapitalMap[rel.no]
	if ! ok {
		return InputOperator()
	}
	rel.char = value

	return rel.char
}