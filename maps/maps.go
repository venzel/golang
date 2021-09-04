package maps

import "fmt"

func Execute(print bool) {

	if print {

		gamers := map[string]uint8{
			"tiago": 30,
			"liz": 5,
		}

		gamers["cintia"] = 20
		
		for k, v := range gamers {
			fmt.Println(k, v)
		}

		winner := gamers["tiago"]

		fmt.Println(winner)

		/* */

		players := map[string]map[string]uint32{
			"tiago": {
				"rounds": 7,
				"victories": 3,
				"defeats": 4,
			},
			"liz": {
				"rounds": 7,
				"victories": 6,
				"defeats": 1,
			},
		}

		players["cintia"] = map[string]uint32{"victories": 12}
		players["venzel"] = map[string]uint32{"victories": 4}

		delete(players, "cintia")

		delete(players, "cintia")

		for k, v := range players {
			fmt.Println(k)
			for i, j := range v {
				fmt.Println(i, j)
			}
		}
	}
}