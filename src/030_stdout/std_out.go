// Использование стандартного потока вывода
package main

import (
	"io"
	"os"
)

// $ go run stdOUT.go
// Please give me one argument!
// $ go run stdOUT.go 123 12
// 123
func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	// io.WriteString() отправляет содержимое своего второго параметра на экран, если первым параметром является os.Stdout
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
