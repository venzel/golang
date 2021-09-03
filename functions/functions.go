package functions

import "fmt"

func Execute(print bool) {

	if print {

		/* */

		var sum = func(a, b int8) int8 {
			return a+b
		}

		fmt.Println(sum(-128, 127))

		/* */

		var gamer = func(name string) (string, int8) {
			return "tiago", 30
		}

		name, age := gamer("tiago")

		fmt.Println(name)
		fmt.Println(age)

		/* */

		var people = func() string {
			var friend = func() string { return "amanda" }

			return friend()
		}

		fmt.Println(people())
	}
}