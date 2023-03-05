package main

import (
	aGamer "appGamer/gamer"
	"fmt"
	"time"
)

func main() {
	addressTiago := aGamer.Address{
		Street: "13 de maio",
		Number: 20,
	}

	tiago := aGamer.Account{
		Name: "Tiago",
		Address: addressTiago,
		Birthday: time.Date(1983, 12, 28, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(tiago)
	fmt.Println(tiago.GetBirthday())
}