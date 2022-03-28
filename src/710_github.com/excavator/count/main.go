// count подсчитывает количество вхождений каждой строки в файле
package main

import (
	"fmt"
	"log"
	"sort"

	"710_github.com/excavator/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}
	// не сортированный вывод
	// for name, count := range counts {
	// 	fmt.Printf("Votes for %s: %d\n", name, count)
	// }
	// сортированный вывод
	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}

}
