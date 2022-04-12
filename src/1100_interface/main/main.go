// тестирование MyInterface
package main

import (
	"1100_interface/define"
	"fmt"
)

func main() {
	var value define.MyInterface
	value = define.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(3.03)
	fmt.Println(value.MethodWithReturnValue())
}
