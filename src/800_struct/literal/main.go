// литерал структуры
package main

import "fmt"

type car struct {
	name     string
	topSpeed int
}

func main() {
	myCar := car{name: "Volga", topSpeed: 123}
	fmt.Println("Name:", myCar.name)
	fmt.Println("Top Speed:", myCar.topSpeed)
}
