// изменение структуры в функции
package main

import "fmt"

type subcriber struct {
	name   string
	rate   float64
	active bool
}

// для изменения структуры передаем указатель
func applyDiscount(s *subcriber) {
	s.rate = 4.99
}

func main() {
	var s subcriber
	//пeредается указатель
	applyDiscount(&s)
	fmt.Println(s.rate)
}
