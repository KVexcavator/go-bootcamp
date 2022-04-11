// инкапсуляция: механизм защиты полей структурного типа от недействительных данных
// сокрытия данных в одной части программы от кода в другой части
// как встраивать другие типы в структуры
package main

import (
	"1000_embed/calendar"
	"fmt"
	"log"
)

func main() {
	// date := Date{Year: 1971, Month: 10, Day: 19} без set методов, нет проверки
	date := calendar.Date{}
	err := date.SetYear(1971)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(19)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	// при наличии гетера
	fmt.Println(date.Year())

	// Date встрено в Event
	event := calendar.Event{}
	err = event.SetTitle("It's today")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())

	// event.month = 5 не отработает
	// event.Date.year = 2019 тоже не работает
	// Экспортируемые методы повышаются так же, как и поля
	err = event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
