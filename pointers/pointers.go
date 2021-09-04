package pointers

import "fmt"

func Execute(print bool) {
	
	if print {

		/* */

		var k int32 = 10
	
		var p *int32 = &k

		fmt.Println(*p)

		/* */

		var x string = "bola"

		var px *string = &x

		fmt.Println(*px)

		/* */

		for i := range x {
			fmt.Println(string(x[i]))
		}

		/* */

		inverter := func(a *int8) {
			*a = *a * -1
		}

		value := int8(10)

		inverter(&value)

		fmt.Println(value)

		/* */

		swap := func(a, b *uint8) {
			aux := uint8(*b)
			*b = *a
			*a = aux
		}

		a := uint8(9)
		b := uint8(8)

		swap(&a, &b)

		fmt.Println(a)
		fmt.Println(b)
	}
}