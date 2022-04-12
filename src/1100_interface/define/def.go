// Определение типа интерфейса состоит из ключевого слова interface, за ним следуют фигурные скобки со списком имен методов, а также параметрами и возвращаемыми значениями, которые должны иметь эти методы.
package define

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

// тип поддерживающий интерфейс
type MyType int

// обязательные методы
func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}
func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

// свои методы
func (my MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
