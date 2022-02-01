// pass_fail сообщает, здал ли пользователь экзамен
package main

import (
	"bufio"   // ввод\вывод с терминала
	"fmt"     // форматирование вывода
	"log"     // обработка ошибок
	"os"      // системные команды
	"strconv" // коныертирование строк
	"strings" // преобразование строк
)

func main() {
	fmt.Print("Inter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64) // 64 всегда передавать и н епариться
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
