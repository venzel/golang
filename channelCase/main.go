package main

import (
	"fmt"
	"time"
)

func work(c chan int, worker string, m map[string]int) {
	for res := range c {
		m[worker] = res
		fmt.Println(worker, res)
		time.Sleep(time.Second)
	}

}

func main() {
	c := make(chan int)

	m := make(map[string]int)

	for i := 0; i < 10; i++ {
		go work(c, fmt.Sprintf("worker%d", i+1), m)
	}

	for i := 0; i < 100; i++ {
		c <- i
	}

	fmt.Println(m)

	close(c)
}
