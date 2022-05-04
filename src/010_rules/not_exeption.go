// как избежать выброса ошибки при подключении неисползуемого пакета
package main

// поставить "_" перед названием в импорте
import (
	"fmt"
	_ "os"
)

func main() {
	fmt.Println("Hello there!")
}
