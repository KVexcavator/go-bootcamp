// получает любое количество аргументов float64 и возвращает наибольшее значение
package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	// отрицательная бесконечность
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
}
