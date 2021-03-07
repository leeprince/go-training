package ctrlconsole

import (
	"fmt"
	"strings"
)

var countryCapitalMap = map[int]string{
	1: "+",
	2: "-",
	3: "*",
	4: "/"}

func InputOperator() string {
	fmt.Printf("【操作符号对应的序号：")

	// map 是无序集合；通过排好序的数组配合集合展示有序数组（数组定长；切片是数组的升级版，自动扩容）
	sortNoSlice := [4]int{1, 2, 3, 4}
	var tips string
	for _, v := range sortNoSlice{
		vv := countryCapitalMap[v]
		tips = fmt.Sprintf("%s %d(%s),", tips, v, vv)
	}
	tips = strings.Trim(tips, ",")
	fmt.Print(tips + "】请输入序号>")

	var no int
	fmt.Scan(&no)

	value, ok := countryCapitalMap[no]
	if ! ok {
		fmt.Println("> > 序号错误！")
		return InputOperator()
	}

	return value
}