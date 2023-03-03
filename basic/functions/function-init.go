package main

import (
	"fmt"
)

func main() {
	fmt.Println("exec 0");
}

// Executa primeiro
fun init() {
	fmt.Println("exec 1");
}

// Executa segundo
fun init() {
	fmt.Println("exec 2");
}