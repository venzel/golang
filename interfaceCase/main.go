package main

import "fmt"

func main() {
    var list []interface{}

    list = append(list, "bola")
    list = append(list, 10)
    list = append(list, true)

    for _, value := range list {
        v, ok := value.(string)

        if ok {
            fmt.Println(v, "is string!")
        }
    }
}