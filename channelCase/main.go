package main

import (
	"fmt"
	"time"
)

func worker(c chan string, gamer *string) {
	for i := 0; i < 10; i++ {
		c <- *gamer
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	c := make(chan string)

	p := make(map[int]string)

	p[0] = "tiago"
	p[1] = "marcos"
	p[2] = "matheus"

	p1, p2, p3 := p[0], p[1], p[2]

	for i := 0; i < 3; i++ {
		go worker(c, &p1)
		go worker(c, &p2)
		go worker(c, &p3)
	}

	for i := 0; i < 10; i++ {
		msg := <-c
		fmt.Println(i, msg)
	}
}
