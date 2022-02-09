package main

import (
	"fmt"
	"reflect"
)

func main(){

	var myBool bool
	// с & получит указатель на адресс в памяти и выведет тип *bool
	fmt.Println(reflect.TypeOf(&myBool))

	// полная запись присваивания
	var myVar int
	// обьявление переменной содержащей указатель на int
	var myVarPointer *int
	// указатель приспаивается переменной
	myVarPointer = &myVar
	// выводится адресс в памяти
	fmt.Println(myVarPointer)


	myInt := 4
	fmt.Println(myInt)
	// короткое обьявление переменной указателя
	myIntPointer := &myInt
	// новое значение для переменной на которую ссылается указатель
	*myIntPointer = 8
	// выводися сам указатель на адресс в памяти
	fmt.Println(myIntPointer)
	// выводится значения переменной (4)
	fmt.Println(*myIntPointer)

}
