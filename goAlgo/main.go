package main

import (
	"fmt"
	"reflect"
)

func tri_a(x []int) []int {
	for i := range x {
		x[i] *= 3
	}
	return x
}

func tri_b(x ...int) []int {
	for i := range x {
		x[i] *= 3
	}
	return x
}

func tri_c(x []interface{}) []interface{} {
	for i := range x {
		if reflect.TypeOf(x[i]).Kind() == reflect.Int {
			x[i] = x[i].(int) * 3
		}
	}
	return x
}

func main() {
	x := []int{1, 2, 3}
	result := tri_a(x)
	fmt.Println(result)

	x1, x2, x3 := 1, 2, 3
	result = tri_b(x1, x2, x3)
	fmt.Println(result)

	k := []interface{}{1, "a", 3}
	resultk := tri_c(k)
	fmt.Println(resultk)

	ix := []interface{}{1, 2, "bola"}
	ix[0] = ix[0].(int) * 3
	ix[1] = ix[1].(int) * 3

}
