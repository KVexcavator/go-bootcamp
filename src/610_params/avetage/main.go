// программа вычисляет среднее от переданных аргументов
package main

import (
	"fmt"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	fmt.Println(average(100, 48))
	fmt.Println(average(43.4, 68.2, 29.3, 123.3))
}
