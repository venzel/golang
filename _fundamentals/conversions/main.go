package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	w := "1"

	value, err := strconv.Atoi(w)

	if err != nil {
		log.Fatal(value)
	}

	fmt.Println(value) // 1

	strVal := "1010"

	vfinal, _ := strconv.ParseInt(strVal, 2, 0)

	fmt.Println(vfinal) // 10

	strValueB := "F"

	vfb, err := strconv.ParseBool(strValueB)

	if err != nil {
		log.Fatal(vfb)
	}

	fmt.Println(vfb) // true

	vlx := 1

	z := strconv.Itoa(vlx)

	fmt.Println(z) // 1

	fmt.Println(reflect.TypeOf(z))
	tp := fmt.Sprintf("%T", z)

	fmt.Println(tp)
}
