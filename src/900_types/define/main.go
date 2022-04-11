// определяемые новые типы на основе базовых
package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	// определение переменных на основе своих типов
	var carFuel Gallons
	// преобразует значение
	carFuel = Gallons(10.00)
	// коороткая запись
	busFuel := Liters(240.00)
	fmt.Println(carFuel, busFuel)

	// значение Liters можно преобразовать в Gallons и наоборот, потому что оба типа имеют базовый тип float64
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
	// определяемый тип поддерживает все те же операции, что и базовый тип.
	fmt.Println(Liters(1.2) < Liters(3.4))
	// определяемый тип может использоваться в операциях вместе с лите­ралами
	fmt.Println(Gallons(1.2) == 1.2)
	// но не могут использоваться в операциях вместе со зна­чениями другого типа
	// fmt.Println(Gallons(1.2) == Liters(1.2)) ошибка

}
