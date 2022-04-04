// шмпортирует структуру из 820_github.com/kvexcavator/magazine/
package main

import (
	"fmt"

	"820_github.com/kvexcavator/magazine"
)

// имя структуры и полей в эспортируемой структуре должны начинаться с большой буквы
func main() {
	// в литерале все поля указывать не обязательно, не указанные будут с нулевым значением
	subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)
}
