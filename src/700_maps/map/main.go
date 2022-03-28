// обьявление карт
// ключ и значение могут быть разных типов, все ключи и все значения должны быть одного типа
package main

import (
	"fmt"
	"sort"
)

func main() {
	// обьявление: клювое слово map , тип ключа, тип значения
	var elements map[string]string
	// создание
	elements = make(map[string]string)
	fmt.Println(elements)

	// обьявление и создание
	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])

	// литерал обьявление со значениями
	myMap := map[string]float64{"a": 1.3, "b": 3.2, "c": 4.2}
	fmt.Println(myMap["a"])
	// for ... range обрабатывает карты в случайном порядке
	for key, value := range myMap {
		fmt.Printf("%s has a value of %0.1f%%\n", key, value)
	}
	for key := range myMap {
		fmt.Println(key)
	}
	for _, value := range myMap {
		fmt.Println(value)
	}
	// применение сортировки
	var keys []string
	for key := range myMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s has a value of %0.1f%%\n", key, myMap[key])
	}

	// проверка, было ли присвоено значение
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok) // true
	// можно игнорировать значение
	_, ok = counters["c"]
	fmt.Println(ok) // false

	// удаление значения с проверкой на присутствие
	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
}
