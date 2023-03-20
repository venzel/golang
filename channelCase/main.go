package main

import "fmt"

func worker(c chan int, list *[]int) {
	for v := range c {
		(*list) = append(*list, v)
		fmt.Println(v)
		// time.Sleep(time.Second)
	}

	close(c)
}

func main() {
	c := make(chan int)

	list := []int{}

	for i := 0; i < 2; i++ {
		go worker(c, &list)
	}

	for i := 1; i < 11; i++ {
		c <- i
	}
}
