// обработка несколких путей
package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}
func rissanHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Privet web!")
}
func kazahHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salam, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/privet", rissanHandler)
	http.HandleFunc("/salam", kazahHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
