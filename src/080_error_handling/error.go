// программа, которая обрабатывает сообщения об ошибках и создает собственное сообщение об ошибке
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		// если аргументы не переданы
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64
	for err != nil {
		if k >= len(arguments) {
			// go run error.go  a b c
			fmt.Println("None of the arguments is a float!")
			return
		}
		// если первый аргумент не float, проверить следующий аргумент
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n
	// определить min и max аргументы
	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}
	// go run error.go b c 1 2 3 c -1 100 -200 a
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
