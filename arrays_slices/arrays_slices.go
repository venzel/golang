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
	}
}