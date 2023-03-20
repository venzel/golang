package main

import "fmt"

func main() {
	channel := make(chan string)

	go write(channel, "bola")
	go write(channel, "bola")

	for {
		message, open := <-channel

		if !open {
			break
		}

		fmt.Println(message)
	}

}

func write(channel chan string, text string) {
	channel <- text

	close(channel)
}
