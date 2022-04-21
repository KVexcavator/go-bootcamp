// рекурсивный вывод всего содержания
package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// обеспечивает продолжение программы после паники
func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		// если значение паники не является признаком ошибки, возобновить панику с темже значением, иначе функция перехватывает любую панику, даже если она не происходит из scanDirectory
		panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// значение ошибки передается панике и возвращать его из функции не надо
		panic(err)
	}

	for _, file := range files {
		// Соединяет путь каталога и имя файла через символ /.
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	defer reportPanic()
	// panic("some other issue") другая паника обрабатывается случаем panic(p)
	scanDirectory("/")
}
