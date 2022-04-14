// Значение ошибки — это любое значение с методом Error, который возвращает строк
package main

import (
	"fmt"
	"log"
)

//Тип error — это всего лишь интерфейс!
type error interface {
	Error() string
}

type ComedyError string

// включение поддержки интерыейса error
func (c ComedyError) Error() string {
	return string(c)
}

// включение поддержки через передачу значения с типом error
type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error

	err = ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)

	err = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
