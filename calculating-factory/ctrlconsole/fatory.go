package ctrlconsole

type OperationFactory interface {
	Calculate(l float64, r float64) float64
}

var factory map[string]OperationFactory

type OperatorAdd struct {}
type OperatorSub struct {}
type OperatorMul struct {}
type OperatorDiv struct {}

// 初始化。默认：引入包时init函数就会执行
func init() {
	factory = make(map[string]OperationFactory)

	factory["+"] = &OperatorAdd{}
	factory["-"] = &OperatorSub{}
	factory["*"] = &OperatorMul{}
	factory["/"] = &OperatorDiv{}
}

// 返回实现OperationFactory接口的方法
func NewOperationFactory(op string) OperationFactory {
	return factory[op]
}

// OperatorAdd 结构体方法实现 OperationFactory 接口
func (data *OperatorAdd) Calculate(l float64, r float64) float64 {
	return l + r
}
// OperatorSub 结构体方法实现 OperationFactory 接口
func (data *OperatorSub) Calculate(l float64, r float64) float64 {
	return l - r
}
// OperatorMul 结构体方法实现 OperationFactory 接口
func (data *OperatorMul) Calculate(l float64, r float64) float64 {
	return l * r
}
// OperatorDiv 结构体方法实现 OperationFactory 接口
func (data *OperatorDiv) Calculate(l float64, r float64) float64 {
	return l / r
}


