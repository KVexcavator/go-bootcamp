// простое ввеб приложение
package main

import (
	"log"
	"net/http"
)

// writer - значение для обновления ответа
// request - запрос от браузера
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello Web!")
	// добавляет строку в ответ
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// если был получен запрос с /hello
	http.HandleFunc("/hello", viewHandler)
	// обработка запроса браузера и ответ на него
	// http://localhost:8080/hello
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
