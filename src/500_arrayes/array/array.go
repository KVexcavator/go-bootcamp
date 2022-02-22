// программа объявляет пару массивов и выводит все их элементы
package main

import (
	"fmt"
)

func main() {
	// массив в go содержит элементы одного типа, по умолчанию элементы имеют нулевое значение для заданного типа
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	// инкремент на единицу
	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	// литерал массива
	var letters = [3]string{"a", "b", "c"}
	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println(letters[0])
	fmt.Println(primes[3])
	fmt.Printf("%#v\n", numbers)

	// Обращение к элементам массива в цикле
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println(index, notes[index])
	index = 3
	fmt.Println(index, notes[index])
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
	// вывод длины массива
	fmt.Println(len(notes))
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	// Безопасный перебор массивов в цикле «for...range»
	for index, note := range notes {
		fmt.Println(index, note)
	}
	// Пустой идентификатор в циклах «for...range»
	for _, note := range notes {
		fmt.Println(note)
	}
	// или
	for index, _ := range notes {
		fmt.Println(index)
	}
}
