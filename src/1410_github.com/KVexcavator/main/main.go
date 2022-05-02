// Ф-я выводит преобразованную строку
package main

import (
	"fmt"

	"1410_github.com/KVexcavator/prose"
)

func main() {
	phrases := []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
