package main

import (
	"module/arrays_slices"
	// "module/examples/player"
	"module/examples/sum"

	// "module/examples/user_register"
	"module/conditional"
	"module/functions"
	"module/loops"
	"module/maps"
	"module/pointers"
	"module/recursion"
	"module/structs"
	"module/variables"
)

func main()  {
	variables.Execute(false)
	loops.Execute(false)
	pointers.Execute(true)
	structs.Execute(false)
	functions.Execute(false)
	arrays_slices.Execute(false)
	conditional.Execute(false)
	recursion.Execute(false)
	maps.Execute(false)
	// user_register.Register()
	// player.Menu()

	// Examples

	sum.Execute(false)
}