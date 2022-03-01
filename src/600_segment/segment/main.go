// программа илистрирует работу с сегментами
package main

import "fmt"

func main() {
	// обьявление сегмента, по факту не создается
	var notes []string
	// создание сегмента из семи строк
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	// короткая запись без обьявления
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	// len отрабатывает также
	fmt.Println(len(notes))
	fmt.Println(len(primes))

	// литерал сегмента, вызывать make не требуется
	letters := []string{"a", "b", "c"}
	// for и for...range работают с сегментами точно так же, как и с массивами
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}
}
