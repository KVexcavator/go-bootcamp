//
package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	// Если тип объявляет методы с указателями на получателей, при присваивании переменным с типом интерфейса можно будет использовать только указатели на этот тип.
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}
