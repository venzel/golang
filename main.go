package main

import (
	"module/arrays_slices"
	// "module/examples/player"
	"module/examples/sum"

	// "module/examples/user_register"
	"module/conditional"
	"module/functions"
	"module/loops"
	"module/pointers"
	"module/structs"
	"module/variables"
)

func main()  {
	variables.Execute(false)
	loops.Execute(false)
	pointers.Execute(false)
	structs.Execute(false)
	functions.Execute(false)
	arrays_slices.Execute(false)
	conditional.Execute(true)
	// user_register.Register()
	// player.Menu()

	// Examples

	sum.Execute(false)
}