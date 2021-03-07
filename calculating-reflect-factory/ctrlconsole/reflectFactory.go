package ctrlconsole

import (
	"fmt"
	"reflect"
)

type Cal struct {
	LeftNum float64
	Operator string
	RigtNum float64
	Result float64
}

/*
反射三大原则：
	原则一：通过接口类型调用 reflect.ValueOf 反射方法可以转换成反射对象
	原则二：反射对象通过 反射对象.interface() 可以转换成接口类型
	原则三：通过 reflect.Value 反射结构体或者 reflect.* 等其他反射方法可以对反射对象的原值修改
 */
func (data *Cal) GetReflectOfCaculete() {
	// 获取对应的反射对象的接口类型
	reflectObj := GetReflectObj(data.Operator)

	/* 反射原则一
	通过反射方法reflect.ValueOf反射出接口类型对应反射对象
	relValue = ReflectAdd ｜ ReflectSub ｜ ReflectMul ｜ ReflectDiv
	 */
	relValue := reflect.ValueOf(reflectObj)

	/* 反射原则三
	通过反射方法reflect.ValueOf和反射结构体reflect.Value设置接口类型对应的已实现接口的结构体方法的参数
	 */
	arg := []reflect.Value{
		reflect.ValueOf(data.LeftNum),
		reflect.ValueOf(data.RigtNum),
	}
	// 通过reflect.Call调用反射对象获取实现OperationFactory接口的具体方法的结构体
	funcStruct := relValue.Call(arg)[0]
	fmt.Printf("valueOper.Call. %v %T", funcStruct, funcStruct)

	/* 反射原则二
	通过类型断言出实现OperationFactory接口的具体方法的结构体
	interfaceFunc = OperatorAdd ｜ OperatorSub ｜ OperatorMul ｜OperatorDiv
	 */
	interfaceFunc := funcStruct.Interface().(OperationFactory)
	fmt.Printf("valueOper.Call. %v %T", interfaceFunc, interfaceFunc)

	// 执行结构体实现OperationFactory接口的方法
	data.Result = interfaceFunc.Calculate()
}
