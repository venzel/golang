package main

import (
	"fmt"
)

// Executa por último
func main() {
	fmt.Println("exec 0");
}

// Executa primeiro
func init() {
	fmt.Println("exec 1");
}

// Executa segundo
func init() {
	fmt.Println("exec 2");
}