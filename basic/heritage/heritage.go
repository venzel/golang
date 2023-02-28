package heritage

import "fmt"

type people struct {
	name string
	age uint8
}

type gamer struct {
	people
	victories int
	defeats int
}

func Execute(print bool) {

	if print {

		p := people{}
		p.name = "tiago"
		p.age = 10

		gamer_a := gamer{p, 10, 9}

		fmt.Println(gamer_a)

		gamer_b := gamer{}
		gamer_b.name = "liz"
		gamer_b.age = 5
		gamer_b.victories = 10
		gamer_b.defeats = 8

		fmt.Println(gamer_b)
		
		gamer_c := gamer{people{"cintia", 10}, 20, 10}

		fmt.Println(gamer_c)
	}
}