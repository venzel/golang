package conditional

import "fmt"

func Execute(print bool) {

	if print {

		/* */

		number_a := 10
		number_b := 20

		if number_a >= 10 {
			fmt.Println("Igual!")
		}

		if number_a > 5 && number_b <= 20 {
			fmt.Println("ok a!")
		}

		if number_a > 5 || number_b > 30 {
			fmt.Println("ok b!")
		}

		if value := 30; value >= 30 {
			fmt.Println("Igual!")
		}

		status := true

		if status {
			fmt.Println("Status ok!")
		}

		/* */

		option := 2

		switch option {
			case 1:
				fmt.Println("option 1")
			case 2:
				fmt.Println("option 2")
			default:
				fmt.Println("option not found!")
		}

		option = 2

		var day string

		switch option {
			case 1:
				day = "domingo"
			case 2:
				day = "segunda"
			case 3:
				day = "terca"
			default:
				day = "option not found!"
		}

		fmt.Println(day)
			
	}
}