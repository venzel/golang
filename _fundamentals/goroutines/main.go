package main

import (
	"fmt"
	"time"
)

func write(text string) {
    for {
        fmt.Println(text)
        time.Sleep(time.Second)
    }
}

func main() {
    go write("ovo") 
    go write("queijo")
    write("pão")
}