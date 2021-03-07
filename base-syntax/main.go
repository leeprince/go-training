package main

import (
	"fmt"
	"reflect"
)

type reflectTestStruct struct {
	name string
	like string
}

func main()  {
	/* & *
	1. 指针与引用：&:返回变量的指针（内存地址） *：返回指针的原值（指针指向的内容）
	2. 变量的指针的类型 等于 *变量类型
	 */
	/*a := 10 // a  0xc0000ae008 =>> 10
	b := &a // b  0xc0000b2018 =>> 0xc0000ae008
	c := &b // b  0xc0000b2018 =>> 0xc0000ae008
	fmt.Printf("a值:%d &a-pts:%p type:%T &a-type:%T \n", a, &a, a, &a)
	fmt.Printf("b值:%d b-pts:%p &b-pts:%p type:%T \n", *b, b, &b, b)
	fmt.Printf("c值:%d c-pts:%p *c-pts:%p type:%T \n", **c, c, *c, c)*/

	/*
	延迟处理：defer
	 */
	/*defer deferFunc()
	fmt.Println("下面执行结束之后，会自动执行延迟处理（defer ...）")*/

	/* 错误处理
	1. panic() 抛出异常
	2. recover() 捕获异常
	3. 通过 defer + recover 延迟处理实现捕获异常
	 */
	/*defer func() {
		fmt.Println("catch exception:", recover())
	}()
	panic("error -- 001")*/

	/*
	封装 try 方法实现自动捕获异常
	 */
	/*try(func () {
		fmt.Println("program process ... after try catch")
		panic("panic on try throw exception")
	}, func (err interface{}) {
		fmt.Println("try catch:", err)
	})*/

	/* 反射
	反射三大原则：
		原则一：通过接口类型调用 reflect.Value 等反射方法可以转换成反射对象
		原则二：反射对象通过 反射对象.interface() 可以转换成接口类型
		原则三：通过 reflect.Value 反射结构体或者 reflect.* 等其他反射方法可以对反射对象的原值修改
	 */
	// i := 10 // 整型
	// i := "10" // 字符串
	// i := [2]int {1, 2} // 数组
	// i := []int {1, 2} // 切片
	/*var i map[string]string // 集合
	i = map[string]string{"name": "prince", "like": "golang"}*/
	/*i := make(map[string]string)  // 集合
	i["name"] = "prince"
	i["like"]  = "golang"*/
	// i := reflectTestStruct{"prince", "golang"} // 结构体
	i := reflectTestStruct{} // 结构体
	i.name = "prince"
	i.like = "golang"
	fmt.Printf("反射得到的类型：%v；反射得到的值:%v \n", reflect.TypeOf(i), reflect.ValueOf(i))



}

func deferFunc()  {
	fmt.Println("defer process end")
}

func try(f func(), catch func(e interface{}))  {
	defer func() {
		err := recover()
		if err != nil {
			catch(err)
		}
	}()

	f()
}