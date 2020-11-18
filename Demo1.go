package main

import (
	"fmt"
	"reflect"
)

type profDemo struct {
	name string
	id   int
}

type studentDemo struct {
	enrollDate   int
	coursesTaken string
	grade        int
	prof         profDemo
}

func basicReflectionDemo(mystery interface{}) {
	mysteryType := reflect.TypeOf(mystery)
	kind := mysteryType.Kind()
	fmt.Println("The type of the parameter is: ", mysteryType)
	fmt.Println("The kind of that parameter is: ", kind)
}

func makeReflectionMap(number interface{}) interface{} {
	map1 := make(map[string]int)
	map2 := make(map[string]float64)

	type1 := reflect.TypeOf(map1)
	type2 := reflect.TypeOf(map2)

	numberType := reflect.TypeOf(number)
	if numberType == reflect.TypeOf(4) {
		return reflect.MakeMap(type1)
	}
	return reflect.MakeMap(type2)

}

func getsizeReflection(mysteryData interface{}) int {
	mysteryType := reflect.TypeOf(mysteryData)
	kind := mysteryType.Kind()
	mysteryValue := reflect.ValueOf(mysteryData)
	switch kind {
	case reflect.TypeOf(profDemo{}).Kind():
		return mysteryValue.NumField()

	case reflect.TypeOf([]int{1}).Kind():
		return mysteryValue.Len()
	default:
		return 1
	}
}

func structRefletion(data interface{}) {
	structReflect := reflect.ValueOf(data)
	numberOFfields := structReflect.NumField()
	for i := 0; i < numberOFfields; i++ {
		fmt.Printf("fieldNumber %d with value %#v \n", i, structReflect.Field(i))
	}
}

func main() {
	var1 := 4
	basicReflectionDemo(var1)
	var2 := []int{1, 2, 4, 5, 6}
	fmt.Println("============================")
	basicReflectionDemo(var2)
	fmt.Println("============================")
	var3 := profDemo{name: "John Santore",
		id: 12345678}
	basicReflectionDemo(var3)
	var4 := studentDemo{
		enrollDate:   345,
		coursesTaken: "Too many make it stop",
		grade:        4,
		prof:         var3,
	}
	structRefletion(var3)
	fmt.Println("============================")
	structRefletion(var4)
}
