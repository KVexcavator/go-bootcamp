// некоторые функции могут получать значения любого типа
package main

import "fmt"

type Whistle string

func (w Whistle) Noise() {
	fmt.Println(w, "is noise")
}

// объявим тип интерфейса, которому вообще не требуются никакие методы
type Anything interface {
}

// Получает параметр с типом пустого интерфейса.
func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	// своих миетодов у пустого нет, чтобы ывзвать надо переопределить
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.Noise()
	}
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("Toyco Canary"))
}
