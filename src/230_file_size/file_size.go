// прога выводит размер файла
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("readme.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
