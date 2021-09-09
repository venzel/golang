package main

import (
	"fmt"
	"time"
)

func write(text string, channel chan string) {
	for i := 1; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}

func main() {
	channel := make(chan string)

	go write("teste...", channel)

	for message := range channel {
		println(message)
	}

	fmt.Println("End program!")
}