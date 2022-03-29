// использование типов с функциями
package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func main() {
	subscriber1 := defaultSubscriber("Adam Smit")
	subscriber1.rate = 4.78
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Viktor Gugo")
	printInfo(subscriber2)
}
