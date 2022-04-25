// ф-я демонстрирует исползование каналов
package main

import "fmt"

// Значение отправляется по каналу
func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	// Создание нового канала
	myChannel := make(chan string)
	go greeting(myChannel)
	// Канал передается функции, выполняющейся в новой горутине.
	fmt.Println(<-myChannel)
}
