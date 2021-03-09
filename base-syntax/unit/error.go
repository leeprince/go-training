package unit

import "fmt"

func Error() {
	/*
		延迟执行：defer
	*/
	defer deferFunc()
	fmt.Println("下面执行结束之后，会自动延迟执行（defer ...）")

	/* 错误处理
	1. panic() 抛出异常
	2. recover() 捕获异常
	3. 通过 defer 延迟执行 + recover 捕获异常实现上层捕获异常
	*/
	defer func() {
		fmt.Println("catch exception:", recover())
	}()
	panic("error -- 001")

	/*
		封装 try 方法实现自动捕获异常
	*/
	try(func() {
		fmt.Println("program process ... after try catch")
		panic("panic on try throw exception")
	}, func(err interface{}) {
		fmt.Println("try catch:", err)
	})
}

/*
延迟执行（defer）的方法
*/
func deferFunc() {
	fmt.Println("defer process end")
}

/*
封装 try 自动捕获异常
*/
func try(f func(), catch func(e interface{})) {
	defer func() {
		err := recover()
		if err != nil {
			catch(err)
		}
	}()

	f()
}
