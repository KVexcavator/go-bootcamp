// передача данных в стандарный потор ошибок
package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}

// При применении bash(1) можно перенаправить стандартный поток ошибок в файл:
// $ go run stdERR.go 2>/tmp/stdError
// This is Standard output
// $ cat /tmp/stdError
// Please give me one argument!
