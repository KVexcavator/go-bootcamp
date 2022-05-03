// передача функции в качестве аргумента другим функциям
package main

import "fmt"

func sayHi() {
	fmt.Println("Hi!")
}

func sayBye() {
	fmt.Println("Bye!")
}

func twice(myFunction func()) {
	myFunction()
	myFunction()
}

func main() {
	twice(sayHi)
	twice(sayBye)
}
