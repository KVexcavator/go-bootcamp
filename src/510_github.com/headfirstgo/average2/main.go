// программа принимает переданные в консоли параметры
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	// fmt.Println(os.Args)     // [./average2 34.4 34.2 45.2 67.1]
	// fmt.Println(os.Args[1:]) // [34.4 34.2 45.2 67.1]

	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	// При вызове функции с переменным количеством аргументов многоточие (...)
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}
