package unit

import "fmt"

func Arr() {
	// 指针数组
	av := [3]int{1, 2, 3}
	var pa [3]*int
	for i, p := range av {
		fmt.Println(i, "---", p)
		pa[i] = &av[i] // pa[i]为地址
		fmt.Println(i, "---", *(pa[i]))
		fmt.Println(i, "---", pa[i])
	}

	// 数组指针
	as := [3]int{1, 2, 3}
	var ap *[3]int
	ap = &as
	fmt.Println(*ap)
}
