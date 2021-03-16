package unit

import (
	"fmt"
	"reflect"
)

type Animal interface {
	getAge() int
	getType() string
}
type dog struct {
	age int
	tp string
}
type cat struct {
	age int
	tp string
}

type UserInfo struct {
	Id     int
	Name   string
	Amount float64
}

/*
接口是一种类型

一个结构体只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。
 */
func Interface() {
	/*
	接口体验
	 */
	d := &dog{1, "dogg"}
	fmt.Println("dog age:", d.getAge())
	fmt.Println("dog tp:", d.getType())
	c := &cat{2, "catt"}
	fmt.Println("cat age:", c.getAge())
	fmt.Println("cat age:", c.getType())

	/*
	接口作为通用类型
	 */
	var i interface{}
	i = "string"
	fmt.Println("通用类型的值i:", i)
	i = 1
	fmt.Println("通用类型的值i:", i)
	i = UserInfo{Id: 2}
	fmt.Println("通用类型的值i:", i.(UserInfo).Id)

	/*
	只能对interface{}类型的变量使用类型断言（类型查询： 变量.(类型)）
	 */
	var ii interface{}
	ii = 1 // 断言 int
	// ii = 2147483640 // 断言 int
	// ii = 1.0 // 断言 float64
	// ii = dog{3, "dog01"} // 断言 Animal
	// 方法一
	v, ok := ii.(Animal)
	if !ok {
		fmt.Printf("断言 Animal 失败.i 类型：%T \n", ii)
	} else {
		fmt.Println("断言 Animal 成功", v.getAge())
	}
	// 方法二
	switch t := ii.(type) {
	case string:
		fmt.Printf("断言 type string: %v \n", t)
	case int:
		fmt.Printf("断言 type int: %v  \n", t)
	case int16:
		fmt.Printf("断言 type int16: %v  \n", t)
	case int32:
		fmt.Printf("断言 type int32: %v  \n", t)
	case float32:
		fmt.Printf("断言 type float32: %v  \n", t)
	case float64:
		fmt.Printf("断言 type float64: %v  \n", t)
	default:
		fmt.Printf("断言 type .switch-default: %v  \n", t)
	}
	// 方法三
	reflectName := reflect.TypeOf(ii)
	fmt.Printf("reflect TypeOf: %v \n", reflectName)
}

func (c *dog) getAge() int {
	return c.age
}
func (c *dog) getType() string {
	return c.tp
}

func (c *cat) getAge() int {
	return c.age
}
func (c *cat) getType() string {
	return c.tp
}
