package ctrlconsole

import (
	"fmt"
	"strings"
)

// map 是无序集合
var countryCapitalMap = map[int]string{
	1: "+",
	2: "-",
	3: "*",
	4: "/"}

func InputOperator() string {
	fmt.Printf("【操作符号对应的序号：")

	// 通过排好序的数组配合集合展示有序数组（数组定长；切片是数组的升级版，无需指定长度而是会自动扩容）
	sortNoArr := [4]int{1, 2, 3, 4} // 数组的定义方式
	// sortNoArr := []int{1, 2, 3, 4} // 切片的定义方式
	var tips string
	for _, v := range sortNoArr{
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