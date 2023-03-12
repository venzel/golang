package main

import "fmt"

func add(p *[]string, text string) {
	(*p) = append((*p), text)
}

func main() {
	itens := make([]string, 0)

	add(&itens, "queijo")
	add(&itens, "pizza")

	fmt.Println(itens)
}
