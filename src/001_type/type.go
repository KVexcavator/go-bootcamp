package main

import(
	"fmt"
	"reflect"
)

func main(){
	var myInt int = 2

	fmt.Println(reflect.TypeOf(float64(myInt)))
}
