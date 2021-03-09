package unit

import "fmt"

var fc func()

func TypeFunc() {
	/* type func()
	type func()
	*/
	typeFunc(testTypeFunc)
}

func typeFunc(f func()) {
	fmt.Println("type func() 开始")
	f()
	fmt.Println("type func() 结束")
}
func testTypeFunc() {
	fmt.Println("----testTypeFunc---")
}
