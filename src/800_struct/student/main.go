// определяет структуру и выводит инфо
package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func printInfo(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func main() {
	var s student
	s.name = "Barmaley"
	s.grade = 1.00
	printInfo(s)
}
