// panic, если doorNumber содержит недопустимое значение
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерирует случайное число от 1 до 3.
func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
