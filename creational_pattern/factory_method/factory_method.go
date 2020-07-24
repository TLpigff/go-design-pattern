/*
  @Time :       2020/4/16 2:39 下午
  @Author :     lyb
  @File :       factory_method
  @Software:    GoLand
  @Description:
*/
package factoryMethod

// Operator是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是 Operator接口实现的基类，封装公用方法
type OperatorBase struct {
	a,b int
}

func (o OperatorBase) SetA(a int) {
	o.a = a
}

func (o OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory 是生产PlusOperator的工厂类
type PlusOperatorFactory struct {}

// PlusOperator 实际的加法实现
type PlusOperator struct {
	OperatorBase
}

func (PlusOperatorFactory) Create() Operator {
	return  &PlusOperator{
		OperatorBase:OperatorBase{},
	}
}

// 计算 a + b
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperatorFactory 是生产MinusOperator的工厂类
type MinusOperatorFactory struct {}

// MinusOperator 实际的减法实现
type MinusOperator struct {
	OperatorBase
}

func (MinusOperatorFactory) Create() Operator {
	return  &MinusOperator{
		OperatorBase:OperatorBase{},
	}
}

// 计算 a - b
func (o MinusOperator) Result() int {
	return o.a - o.b
}








