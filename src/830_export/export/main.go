// шмпортирует структуру из 820_github.com/kvexcavator/magazine/
package main

import (
	"fmt"

	"820_github.com/kvexcavator/magazine"
)

// имя структуры и полей в эспортируемой структуре должны начинаться с большой буквы
func main() {
	// в литерале все поля указывать не обязательно, не указанные будут с нулевым значением
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)
	fmt.Println(subscriber.HomeAddress)
	fmt.Println("Postal Code:", subscriber.HomeAddress.PostalCode)

	// анонимный вызов
	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("Postal Code:", employee.PostalCode)
}
