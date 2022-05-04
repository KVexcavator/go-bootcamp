// как загружать внешние пакеты
package main

// go get -v github.com/mactsouk/go/simpleGitHub
import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}
