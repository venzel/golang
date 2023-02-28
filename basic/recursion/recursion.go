package recursion

import "fmt"

func fib(position uint) uint {
	if position <= 1 {
		return position
	}

	return fib(position - 2) + fib(position - 1)
}

func Execute(print bool) {

	if print {
		position := uint(10)

		fmt.Println(fib(position))

		for i := uint(1); i <= position; i++ {
			fmt.Println(fib(i))
		}
	}
}