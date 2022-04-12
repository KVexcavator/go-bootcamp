// переменная с типом интерфейса, может принимать значения любого типа, поддерживающего интерфейс
package main

import (
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Robot drank")
}

func (r Robot) Walk() {
	fmt.Println("Robot steps")
}

type NoiseMaker interface {
	MakeSound()
}

// Параметры функций также могут объявляться с типами интерфейсов
func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toy Whistle")
	toy.MakeSound()
	toy = Horn("Toy Horn")
	toy.MakeSound()
	toy = Robot("Robot Bobot")
	//toy.Walk() определен по интерфейсу и не отработает
	// обратное утверждение
	var robot Robot = toy.(Robot)
	robot.Walk()

	play(Whistle("Play whistle"))
	play(Horn("Play Horn"))
}
