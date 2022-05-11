// запись ошибок в кастомный лог файл
package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "tmp/my_log.log"

// $ go run custom_log.go
// $ cat tmp/my_log.log
func main() {
	// 0644 права доступа к файлу
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	// iLog.SetFlags(log.LstdFlags) флаг времени и даты
	iLog.SetFlags(log.LstdFlags | log.Lshortfile) // 2й флаг имя файла и строка
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
