package main

import (
	"fmt"
	"reflect"
)

func basicReflectionDemo(mystery interface{}){
	mysteryType := reflect.TypeOf(mystery)
	kind := mysteryType.Kind()
	fmt.Println("The type of the parameter is: ", mysteryType)
	fmt.Println("The kind of that parameter is: ", kind)
}


func main() {
	var1:=4
	basicReflectionDemo(var1)
	fmt.Println("============================")
}
