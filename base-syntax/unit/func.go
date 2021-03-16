package unit

import "fmt"

var fc func()

type tp func(s string, x, y int) string

/*
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
函数可以返回任意数量的返回值。使用关键字 func 定义函数，左大括号依旧不能另起一行
*/
func Func() {
	/* type func()
	type func()
	*/
	typeFunc(testTypeFunc)

	/*
		tFunc
	*/
	i := tFunc(func() int {
		return 1
	})
	fmt.Printf("tFunc 返回：%v \n", i)

	/*
		tpFunc
	*/
	tpv := tpFunc(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "x:%d;y:%d", 100, 101)
	fmt.Printf("tpFunc 返回：%v \n", tpv)

	/*
	当调用函数，传递过来的变量就是函数的实参，函数可以通过两种方式来传递参数:
	1. 值传递:指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	2. 引用传递:是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

	>
	注意1:无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
	注意2:map、slice、chan、指针、interface默认以引用的方式传递。
	不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。(可变参数)
	Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。
	在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“...”即可
	 */
	// 引用传递
	x := 1
	y := 2
	fmt.Printf("swap 前-x:%v;y:%v \n", x, y)
	swap(&x, &y)
	fmt.Printf("swap 后-x:%v;y:%v \n", x, y)

	/* 返回值
	"_"标识符，用来忽略函数的某个返回值
	Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。返回值的名称应当具有一定的意义，可以作为文档使用。
	没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
	Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
	 */
	s, sq := add(2, 4)
	fmt.Printf("返回值。s: %v; sq: %v \n", s, sq)
	s1, sq1 := addNoReturn(2, 4)
	fmt.Printf("返回值。s1: %v; sq1: %v \n", s1, sq1)

}

func add(x, y int) (sum int, square int) {
	c := x + y
	d := x*x
	return c, d
}

func addNoReturn(x, y int) (sum int, square int) {
	sum = x + y
	square = x*x
	return
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func tFunc(fn func() int) int {
	return fn()
}

func tpFunc(fn tp, s string, x, y int) string {
	return fn(s, x, y)
}

func typeFunc(f func()) {
	fmt.Println("type func() 开始")
	f()
	fmt.Println("type func() 结束")
}
func testTypeFunc() {
	fmt.Println("----testTypeFunc---")
}
