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

	defer file.Close()

	if err != nil {
		panic(err)
	}

	file.Write([]byte("bola"))
}
