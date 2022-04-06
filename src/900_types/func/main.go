//  проеобразование типов с помощью функций
package main

import "fmt"

type Liters float64
type Gallons float64

func toGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func toLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)
	carFuel += toGallons(Liters(40.0))
	busFuel += toLiters(Gallons(100.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
