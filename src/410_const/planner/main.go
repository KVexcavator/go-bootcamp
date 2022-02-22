package main

import (
	"410_const/dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}
