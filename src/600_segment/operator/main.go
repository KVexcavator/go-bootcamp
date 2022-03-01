// программа идюстрирует создание сегмента
// сам сегмент не содержит данных, это всего лишь «окно» для просмотра элементов базового массива
package main

import "fmt"

func main() {
	// сегмент создается на основе базового массива
	// при помощи оператора сегмента
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// вернет a,b,c (третья позиция не включительно)
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)
	// можно передать в переменные
	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2) // b,c,d
	// что бы включить последний элумент надо указать на один больше
	slice3 := underlyingArray[0:5]
	fmt.Println(slice3)
	// если начальный индекс не указан, он будет ноль
	slice4 := underlyingArray[:3]
	fmt.Println(slice4)
	// усли не указан конесный, включаются все
	slice5 := underlyingArray[1:]
	fmt.Println(slice5)

	// из базового массива можно создать, несколко сегментов
	array := [5]string{"a", "b", "c", "d", "e"}
	segment1 := array[0:3]
	segment2 := array[2:5]
	fmt.Println(segment1, segment2)
	// при изменении массива, меняется сегмент
	// изменения элементов массива будут видны во всех сегментах
	segment3 := array[1:]
	array[1] = "X"
	fmt.Println(segment1, segment3)
	// изменение сегмента, изменяет базовый массив
	segment2[1] = "Y"
	fmt.Println(array)
}
