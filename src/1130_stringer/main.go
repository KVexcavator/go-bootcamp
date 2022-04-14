// при выводе все значения выглядят одинаково
package main

import "fmt"

// добавляется описание
type Stringer interface {
	String() string
}

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gallons", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f liters", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f milliliters", m)
}

func main() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
}
