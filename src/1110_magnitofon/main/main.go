// программа исспользует пакет gadget
package main

import (
	"1110_magnitofon/gadget"
	"fmt"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("A botel of Rome")
	player.Stop()
	// утверждение типа для перехода к конткретному типу
	// ok передается для подтверждения иссползования конткретного типа
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		// вызов метода для конкретного типа
		recorder.Record()
	}
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	fmt.Println("----------------")

	TryOut(gadget.TapePlayer{})
	TryOut(gadget.TapeRecorder{})
}
