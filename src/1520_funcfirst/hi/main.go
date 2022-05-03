// ф-я первого класса может принимпть фргументом другую ыункцию
package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}
func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var greeterFunction func()
	// У функций, которые получают функцию в параметре, также должны быть указаны параметры и возвращаемые типы, которыми должна обладать передаваемая функция.
	var mathFunction func(int, int) float64
	greeterFunction = sayHi
	mathFunction = divide
	greeterFunction()
	fmt.Println(mathFunction(5, 2))
}
