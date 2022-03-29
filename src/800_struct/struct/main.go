// обьявление структур
package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 3.14
	fmt.Println(myStruct.number)

	var subscriber struct {
		name   string
		rate   float64
		active bool
	}

	subscriber.name = "Pol Anderson"
	subscriber.rate = 4.2
	subscriber.active = true

	fmt.Println("Name:", subscriber.name)
	fmt.Println("Monthly rate:", subscriber.rate)
	fmt.Println("Active?:", subscriber.active)
}
