// функция получает минимальное значение, максимальное значение и любое количество дополнительных аргументов float64, игнорирует аргументы, меньшие заданного минимума или большие максимума, и возвращает сегмент с аргументами, принадлежащими указанному диапазону
package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
}
