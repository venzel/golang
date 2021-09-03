package structs

import "fmt"

type Gamer struct {
	name string
	age int32
}

func dataGamer(d Gamer) {
	fmt.Println(d.name, d.age)
}

func Execute(print bool) {

	if print {

		/* */

		var g Gamer
		
		g.name = "tiago"
		g.age = 30

		fmt.Println(g)

		/* */

		gamer_a := Gamer{"tiago", 10}

		fmt.Println(gamer_a)

		/* */

		gamer_b := Gamer{name: "tiago"}

		fmt.Println(gamer_b)

		/* */

		gamers := []Gamer{}

		gamer_c := Gamer{name: "Liz"}
		gamer_d := Gamer{"Cintia", 25}

		gamers = append(gamers, gamer_c)
		gamers = append(gamers, gamer_d)
		gamers = append(gamers, Gamer{"venzel", 30})

		fmt.Println(len(gamers))

		for i := range gamers {
			dataGamer(gamers[i])
		}

		/* */

		list := []Gamer{
			{"xitao", 20},
			{"xororo", 20},
			{"tiririca", 20},
		}

		for i := 0; i < len(list); i++ {
			dataGamer(list[i])
		}
	}
}