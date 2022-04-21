// ф-я перебирает директории с файлами
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// запускаю из корневой директории
	files, err := ioutil.ReadDir("src/1210_list/files/my_dir")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
