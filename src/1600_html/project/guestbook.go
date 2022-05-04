// приложение — гостевая книга для веб-сайта
package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Guestbook - структура, используемая при отображении view.html.
type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

// check вызывает log.Fatal для любых ошибок, отличных от nil.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// чтение строк из файла
func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

// обработчик
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	// создаем новую сигнатуру Guestbook
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	// содержимое шаблона записывается в ResponseWriter
	err = html.Execute(writer, guestbook)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

// http.Request содержит данные полученные из формы
func createHandler(writer http.ResponseWriter, request *http.Request) {
	// получает значение поля signature из формы
	signature := request.FormValue("signature")
	// открывается или создается в файл
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	// записывается текст в новой строке файла
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	// перенаправляем и возвращаем статус
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

// вызывает http://localhost:8080/guestbook
func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	// err никогда не равен nil, пожтому check не мспользуется
	log.Fatal(err)
}
