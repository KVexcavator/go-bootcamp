// Вызов из Go C-кода в отдельных файлах
package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.cHello()

	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Mihalis!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)
	fmt.Println("All perfectly done!")
}

// Скомпилировать С библиотеку
// $ ls -l callClib/
// $ gcc -c callClib/*.c
// $ ls -l callC.o
// $ file callC.o
// $ /usr/bin/ar rs callC.a *.o
// $ ls -l callC.a
// $ file callC.a
// $ rm callC.o

///Собрать и вызвать
// $ go build callC.go
// $ ls -l callC
// $ file callC
// $ ./callC
