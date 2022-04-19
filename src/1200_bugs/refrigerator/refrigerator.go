// функция холодильника с продуктами
package main

import (
	"fmt"
	"log"
)

// ф-я поиска строки в сегменте
func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

// в сегменте хранятся названия продуктов
type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening refrigerator")
}

func (r Refrigerator) Close() {
	fmt.Println("Closing refrigerator")
}

func (r Refrigerator) FindFood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}
	return nil
}

func main() {
	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
