// set методы проверяют значение
package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

// нужно передавать указатель
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func main() {
	date := Date{}
	err := date.SetYear(1971)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)
}
