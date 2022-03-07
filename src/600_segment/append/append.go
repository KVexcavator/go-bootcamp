// функция append получает сегмент и значения, которые присоединяются в конец сегмента
package main

import "fmt"

func main() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	// Если присваивать сегменты разным переменным, может возникнуть путаница. Поэтому существует соглашению о присваивании возвращаемого значения той же переменной
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	// здесь начнется бардак
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)

	//при обращении к элементу сегмента, которому не было присвоено значение, вурнется нулевое значение для этого типа
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5]) // 0 false
	// переменная сегмента, которой не был присвоен сегмент, будет содержать значение nil
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice) // intSlice: []int(nil), stringSlice: []string(nil)
	// функция len, если передать ей вместо сегмента nil, вернет 0
	fmt.Println(len(intSlice))
	// при передаче в пустой сегмент, функции «append» вернет сегмент из одного элемента
	if len(intSlice) == 0 {
		intSlice = append(intSlice, 27)
	}
	fmt.Printf("intSlice: %#v\n", intSlice) // stringSlice: []string{27}
}
