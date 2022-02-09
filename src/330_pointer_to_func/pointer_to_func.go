// пример присвоения указателя ыункции
package main

import (
	"fmt"
)

// обьявляем, что функция возвращает указатель
func createPointer() *float64 {
	myFloat := 0.57
	// возвращается указатель заданого типа
	return &myFloat
}

// передача указателя в функции аргументу
func printPointer(myBoolPointer *bool){
	fmt.Println(*myBoolPointer)
}

func main(){
	// назначить заданный указатель переменной
	var myFloatPointer *float64 = createPointer()
	// выводим значение переменной
	fmt.Println(*myFloatPointer)

	// прием указателя из аргумента
	var myBool bool = true
	printPointer(&myBool)
}
