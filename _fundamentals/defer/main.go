package main

import (
	"fmt"
	"os"
)

func execute() {
    fmt.Println("executou!")
}

func main() {
    defer execute()
    
    file, err := os.Create("./gamer.txt")
    
    if err != nil {
        panic(err)
    }
    
    defer file.Close()
    
    file.Write([]byte("bola"))
}