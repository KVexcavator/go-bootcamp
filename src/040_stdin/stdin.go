// Чтение стандартного потока ввода
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Каждая прочитанная строка выводится на экран, после чего считывается следующая строка
func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	// bufio.NewScanner() возвращает переменную bufio.Scanner , которая используется функцией Scan() для построчного чтения из os.Stdin
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
