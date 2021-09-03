package variables

import (
	"errors"
	"fmt"
)

func Execute(print bool) {

	if print {
		name := "tiago"
		age := 10
		allowed := true;
		activated, enabled := true, true
		price := 31.32
		letter := 'c'

		var max byte = 255
		var pos uint32 = 30
		var posMax uint8 = 255
		var erro error = errors.New("tecnical error")

		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(allowed)
		fmt.Println(price)
		fmt.Println(activated)
		fmt.Println(enabled)
		fmt.Println(max)
		fmt.Println(pos)
		fmt.Println(posMax)
		fmt.Println(string(letter))
		fmt.Println(erro)
	}
}