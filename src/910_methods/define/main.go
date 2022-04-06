// определение метода, в Go метод не тоже самое что и функция
// у метода перед име­нем функции добавляется один дополнительный параметр, называемый параметром получателя
package main

import "fmt"

type MyType string

// у каждого созданного типа могут быть свои методы, похоже на сингклтон в ruby
// метод экспортируется из текущего пакета, если его имя начинается с буквы верхнего регистра
func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

//  можно передавать параметры в методы
func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

// для метода можно объявить одно или несколько возвращаемых значений
func (m MyType) WithReturn() int {
	return len(m)
}

func main() {
	value := MyType("a myType value")
	value.sayHi()
	anotherValue := MyType("a another value")
	anotherValue.sayHi()
	value.MethodWithParameters(4, true)
	fmt.Println(value.WithReturn())
}
