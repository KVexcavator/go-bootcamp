// как использовать хеш-таблицу для хранения всех указателей в виде целых чисел
package main

import "runtime"

func main() {
	var N = 40000000
	// myMap := make(map[int]int)  без ууказателей
	myMap := make(map[int]*int) // с указателями
	for i := 0; i < N; i++ {
		value := int(i)
		// myMap[value] = value без указателей
		myMap[value] = &value // с указателем
	}
	runtime.GC()
	_ = myMap[0]
}
