package unit

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type reflectTestStruct struct {
	Name string
	Like string
}

func Reflect() {
	/* 反射
	反射三大原则：
		原则一：通过接口类型调用 reflect.Value 等反射方法可以转换成反射对象
		原则二：反射对象通过 反射对象.interface() 等方法可以转换成接口类型
		原则三：通过 reflect.Value 等其他反射方法可以对反射对象的原值修改
	*/
	// i := 10 // 整型
	// i := "10" // 字符串
	// i := [2]int{1, 2} // 数组
	// i := []int{1, 2} // 切片
	/*var i map[string]string // 集合
	i = map[string]string{"name": "prince", "like": "golang"}*/
	/*i := make(map[string]string)  // 集合
	i["name"] = "prince"
	i["like"]  = "golang"*/
	// i := map[string]string{"name": "prince", "like": "golang"} // 集合[推荐]
	/*i := reflectTestStruct{} // 结构体
	i.name = "prince"
	i.like = "golang"*/
	// var i = reflectTestStruct{Name: "prince", Like: "golang"} // 结构体
	// i := reflectTestStruct{Name: "prince", Like: "golang"} // 结构体
	i := reflectTestStruct{"prince", "golang"} // 结构体[推荐]
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
		fmt.Println("---反射获取字段---")
		fmt.Println("反射获取的字段:", rt.Field(i))
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
	/* 设置结构体数据的原值
	当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)
	*/
	rvv := reflect.ValueOf(&i)               // 必须是引用类型（变量指针）
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
带标签的结构体
*/
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func ReflectStruct() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println("------------------")
	fmt.Printf("获取到的字段：%v \n", ixField)
	fmt.Printf("获取到结构体字段的名称：%v \n", ixField.Name)
	fmt.Printf("获取到结构体字段的类型：%v \n", ixField.Type)
	fmt.Printf("获取到结构体字段的标签：%v \n", ixField.Tag)
}

/*
 反射的 Type 与 反射的 Value
	变量的最基本信息就是类型和值：反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。
*/
func ReflectTypeAndValue() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}

/*
反射与Printf
*/
type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) { // type switch
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		// more types
		default:
			os.Stdout.WriteString("???")
		}
	}
}
func ReflectPrintf() {
	print(Day(1), "was", Celsius(18.36))
}
