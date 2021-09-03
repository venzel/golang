package loops

import "fmt"

func Execute(print bool) {

	if print {

		/* */

		for i := 0; i < 4; i++ {
			fmt.Println(i)
		}

		/* */

		k := 0

		for {
			k++;

			fmt.Println(k)

			if k == 5 {
				break
			}
		}

		/* */

		gamers := []string{"Tiago"}

		gamers = append(gamers, "Liz")

		for i := range gamers {
			println(gamers[i])
		}
	}
}