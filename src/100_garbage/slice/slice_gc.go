// использования среза и сборщика мусора
package main

import "runtime"

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	runtime.GC()
	// для предотвращения преждевременной очистки сборщиком мусора переменной structure
	_ = structure[0]
}
