// average вычисляет среднее значение.
// go install 510_github.com/headfirstgo/average
// average появилось в папке bin
package main

import (
	"fmt"
	"log"

	"510_github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("../data/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
