package main

import (
	"fmt"
	"reflect"
)

type reflectTestStruct struct {
	Name string
	Like string
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
	延迟执行：defer
	 */
	/*defer deferFunc()
	fmt.Println("下面执行结束之后，会自动延迟执行（defer ...）")*/

	/* 错误处理
	1. panic() 抛出异常
	2. recover() 捕获异常
	3. 通过 defer 延迟执行 + recover 捕获异常实现上层捕获异常
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
		原则二：反射对象通过 反射对象.interface() 等方法可以转换成接口类型
		原则三：通过 reflect.Value 等其他反射方法可以对反射对象的原值修改
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
	/*i := reflectTestStruct{} // 结构体
	i.name = "prince"
	i.like = "golang"*/
	// i := reflectTestStruct{"prince", "golang"} // 结构体
	i := reflectTestStruct{Name: "prince", Like: "golang"} // 结构体
	/*
	反射第一原则：通过接口类型调用 reflect.Value 等反射方法可以转换成反射对象
	 */
	rt := reflect.TypeOf(i)
	rv := reflect.ValueOf(i)
	fmt.Printf("反射得到的类型：%v；反射得到的值: %v \n", rt, rv)
	fmt.Printf("反射得到的类型的数据类型：%T；反射得到的值的数据类型: %T \n", rt, rv)
	/* 结构体为变量测试获取字段信息。注：反射实例 i 的数据类型必须是结构体才拥有NumField方法，否则抛出异常 */
	rtn := rv.NumField() // 获取变量包含字段总数。
	fmt.Println("获取变量包含字段总数: ", rtn)
	for i := 0; i < rtn; i++ {
		fmt.Println("反射获取的字段名称:", rt.Field(i).Name)
	}
	/*
	反射第二原则：反射对象通过 反射对象.interface() 等方法可以转换成接口类型
	 */
	rInterface := rv.Interface()
	fmt.Printf("反射对象.interface() 转换成的接口类型: %T \n", rInterface)
	typeAssertion, ok := rInterface.(reflectTestStruct) // 断言(类型断言:类型判断): 返回原值
	if ok {
		fmt.Println("断言成功，数据类型原值为:", typeAssertion)
	} else { // ttype.(int)、ttype.(string) ...
		fmt.Println("断言成功失败")
	}
	interfaceType := rt.Name() // 获取接口类型的数据类型
	fmt.Println("获取接口类型的数据类型:", interfaceType)
	/*
	反射原则三：通过 reflect.Value 等其他反射方法可以对反射对象的原值修改
	 */
	/* 设置结构体数据的原值 */
	rvv := reflect.ValueOf(&i)
	vField := rvv.Elem().FieldByName("Name") // 获取指定字段。注意：（golang 的大小写：public： 公共的、外部包可访问。proteced: 仅同包可访问）结构体不同包首字母必须大写，相同包此处的字段名也必须是大写
	vField.SetString("prince01")
	fmt.Println("通过 reflect.Value 等其他反射方法可以对反射对象的原值修改结果：", i)
	vvField := rvv.Elem().FieldByName("Like") // 获取指定字段，并且该字段是可访问的（必须是大写开头）
	vvField.SetString("golnag01")
	fmt.Println("通过 reflect.Value 等其他反射方法可以对反射对象的原值修改结果：", i)
	/* 设置整型数据的原值 */
	ii := 10
	rvii := reflect.ValueOf(&ii) // 必须是引用类型（变量指针）
	rvii.Elem().SetInt(100)
	fmt.Println("通过 reflect.Value 等其他反射方法可以对反射对象的原值修改结果：", ii)
	/* 设置字符串数据的原值 */
	ss := "princeValue"
	rvs := reflect.ValueOf(&ss) // 必须是引用类型（变量指针）
	rvs.Elem().SetString("princeValue01")
	fmt.Println("通过 reflect.Value 等其他反射方法可以对反射对象的原值修改结果：", ss)






}

/*
延迟执行（defer）的方法
 */
func deferFunc()  {
	fmt.Println("defer process end")
}

/*
封装 try 自动捕获异常
 */
func try(f func(), catch func(e interface{}))  {
	defer func() {
		err := recover()
		if err != nil {
			catch(err)
		}
	}()

	f()
}