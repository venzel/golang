package functions

import (
	"fmt"
)

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

		/* */

		var sm = func (title string, numbers ...int) int {
			fmt.Println(title)

			total := 0

			for i := range numbers {
				total += numbers[i]
			}

			return total
		}

		fmt.Println(sm("soma", 3, 4, 9, 21))

		/* */

		var soma = func(userName string, values ...uint8) (name string, count uint8) {
			name = userName;
			count = 0;
	
			for i := range values {
				count += values[i];
			}
	
			return;
		}
	
		name, count := soma("tiago", 2, 4, 4);
	
		fmt.Println(name, count);
	}
}