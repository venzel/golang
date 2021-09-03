package arrays_slices

import "fmt"

func Execute(print bool) {

	if print {

		/* */

		var names [3]string

		names[0] = "tiago"
		names[1] = "liz"
		names[2] = "cintia"

		var i uint8 = 0
		var f uint8 = 2

		fmt.Println(names[i:f])

		/* */

		gamers := []string {"tiago", "liz"}

		gamers = append(gamers, "cintia")

		fmt.Println(gamers)

		/* */

		i = 1
		f = 2

		gamers_filter := gamers[i:f]

		fmt.Println(gamers_filter)

		/* */

		arr := make([]float32, 2, 4)

		arr = append(arr, 3.1)
		arr = append(arr, 3.2)
		arr = append(arr, 3.3)

		fmt.Println(arr)
		fmt.Println(len(arr))
		fmt.Println(cap(arr))
	}
}