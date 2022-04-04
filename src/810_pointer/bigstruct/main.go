// указатели гарантируют, что в памяти будет храниться только одна копия каждой структуры
package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

// для здоровья памяти ставить везде указатели
func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	// это не структура а указатель
	subscriber1 := defaultSubscriber("Aman Singh")
	// & не нужен , так как это структура
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}
