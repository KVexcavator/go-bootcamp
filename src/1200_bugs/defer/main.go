// Если у вас есть вызов функции, который должен быть гарантированно выполнен в любом случае, используйте команду defer.
package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	// defer вызовется последней, но гарантированно
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	// вернет ошибку после defer
	return fmt.Errorf("I don't want to talk.")
	// не будет исполнено
	fmt.Println("Nice weather, eh?")
	return nil
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
