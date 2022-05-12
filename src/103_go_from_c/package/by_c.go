// вызов go-пакета из С-программы
package main

import "C"
import (
	"fmt"
)

// export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

// export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {}

// сгенерировать из Go-кода общую библиотеку C
// $ go build -o by_c.o -buildmode=c-shared by_c.go // должен создать два файла .c и .h
// $ file by_c.o

// $ gcc -o willUseGo willUseGo.c ./by_c.o
// $ ./willUseGo
