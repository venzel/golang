package sum

import "fmt"

func sum(values []int32) int32 {
	var amount int32 = 0

	for _, v := range values {
		amount += v
	}

	return amount
}

func Execute(print bool) {
	
	if print {
		numbers := []int32{20, 30, 30, 10, 2}

		amount := sum(numbers)

		fmt.Println(amount)
	}
}