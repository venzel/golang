package methods

import "fmt"

var age uint8

type gamer struct {
	name string
	victories uint8
}

func (g *gamer) save(name string) {
	g.name = name
	g.victories = 8
}

func (g gamer) approved() bool {
	return g.victories > 10
}

func init() {
	age = uint8(20)
}

func Execute(print bool) {

	if print {
		fmt.Println(age)

		gamer := gamer{}

		gamer.save("marcos")

		fmt.Println(gamer)

		fmt.Println(gamer.approved())
	}
}