// ф-я демонстрирует синхронизацию горутин
package main

import "fmt"

func even(channel chan string) {
	channel <- " 1_even "
	channel <- " 2_even "
	channel <- " 3_even "
}

func odd(channel chan string) {
	channel <- " 1_odd "
	channel <- " 2_odd "
	channel <- " 3_odd "
}

// Выводит 1_even  1_odd  2_even  2_odd  3_even  3_odd
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go even(channel1)
	go odd(channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}
