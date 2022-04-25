// ф-я демонстрации горутины
package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}
func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}
func main() {
	// добавить go для запуска конкурентного выполнения
	go a()
	go b()
	// выполнение main может опережать горутины, и выведется только "end main()"для наглядности можно поспать
	time.Sleep(time.Second)
	fmt.Println(" end main()")
}
