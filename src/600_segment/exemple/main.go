package main

import "fmt"

func main() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 17.2
	for i, nunumber := range numbers {
		fmt.Println(i, nunumber)
	}

	var letters = []string{"a", "b", "c"}
	for i, leletter := range letters {
		fmt.Println(i, leletter)
	}
}
