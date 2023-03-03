package function_init

import (
	"fmt"
)

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