// определение собственного типа и структуры
package main

import "fmt"

// ключевое слово type, определение типа с именем part
type part struct {
	description string
	count       int
}

// переменная с типом car
type car struct {
	name     string
	topSpeed float64
}

// передача своего типа в параметрах функции
func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

// функция сщздания part
func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	// обьявление переменной типа сar
	var porsche car
	porsche.name = "Porsche 931 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)

	p := minimumOrder("Small Hex")
	fmt.Println(p.description, p.count)
}
