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
	}
}