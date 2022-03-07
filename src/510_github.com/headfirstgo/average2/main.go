// программа принимает переданные в консоли параметры
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(os.Args)     // [./average2 34.4 34.2 45.2 67.1]
	// fmt.Println(os.Args[1:]) // [34.4 34.2 45.2 67.1]

	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
