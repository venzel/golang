package main

import "fmt"

func reverse[T interface{}](list []T) []T {
    aux := make([]T, len(list))

    count := len(list) - 1;

    for i := range list {
        aux[i] = list[count]
        count--
    }

    return aux
}

func reversex[T string | int](list *[]T) []T {
    aux := make([]T, len(*list))

    count := len(*list) - 1;

    for i := range *list {
        aux[i] = (*list)[count]
        count--
    }

    return aux
}

func main() {
    letters := []string{"a", "b", "c"}

    numbers := []int{8, 9, 10}

    fmt.Println(reverse(letters))

    fmt.Println(reversex(&numbers))
}