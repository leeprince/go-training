package ctrlconsole

// 接口
type OperationFactory interface {
	Calculate() float64
}

// 计算的数据结构体
type caculateNum struct {
	l float64
	r float64
}

/*
引用caculateNum计算的数据结构体
 */
type OperatorAdd struct {
	num *caculateNum
}
type OperatorSub struct {
	num *caculateNum
}
type OperatorMul struct {
	num *caculateNum
}
type OperatorDiv struct {
	num *caculateNum
}

// 反射对象map的映射
var factory map[string]interface{}

// 初始化。默认：引入包时init函数就会执行
func init() {
	factory = make(map[string]interface{})
	factory["+"] = ReflectAdd
	factory["-"] = ReflectSub
	factory["*"] = ReflectMul
	factory["/"] = ReflectDiv
}

// 获取对应的反射对象的接口类型
func GetReflectObj(op string) interface{} {
	return factory[op]
}

/*
反射对象
 */
// 反射对象的接口类型对应的反射对象（该反射对象返回实现OperationFactory接口的具体方法的结构体）
func ReflectAdd(l float64, r float64) *OperatorAdd {
	return &OperatorAdd{
		num: &caculateNum{
			l: l,
			r: r,
		},
	}
}
// 反射对象的接口类型对应的反射对象（该反射对象返回实现OperationFactory接口的具体方法的结构体）
func ReflectSub(l float64, r float64) *OperatorSub {
	return &OperatorSub{
		num: &caculateNum{
			l: l,
			r: r,
		},
	}
}
// 反射对象的接口类型对应的反射对象（该反射对象返回实现OperationFactory接口的具体方法的结构体）
func ReflectMul(l float64, r float64) *OperatorMul {
	return &OperatorMul{
		num: &caculateNum{
			l: l,
			r: r,
		},
	}
}
// 反射对象的接口类型对应的反射对象（该反射对象返回实现OperationFactory接口的具体方法的结构体）
func ReflectDiv(l float64, r float64) *OperatorDiv {
	return &OperatorDiv{
		num: &caculateNum{
			l: l,
			r: r,
		},
	}
}

/*
接口的实现
 */
// OperatorAdd 结构体方法实现 OperationFactory 接口
func (data *OperatorAdd) Calculate() float64 {
	return data.num.l + data.num.r
}
// OperatorSub 结构体方法实现 OperationFactory 接口
func (data *OperatorSub) Calculate() float64 {
	return data.num.l - data.num.r
}
// OperatorMul 结构体方法实现 OperationFactory 接口
func (data *OperatorMul) Calculate() float64 {
	return data.num.l * data.num.r
}
// OperatorDiv 结构体方法实现 OperationFactory 接口
func (data *OperatorDiv) Calculate() float64 {
	return data.num.l / data.num.r
}
