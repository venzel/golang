package main

import (
	"module/arrays_slices"
	"module/conditional"
	"module/examples/send_mail"
	"module/examples/sum"
	"module/functions"
	"module/interfaces"
	"module/loops"
	"module/maps"
	"module/methods"
	"module/pointers"
	"module/recursion"
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
	conditional.Execute(false)
	recursion.Execute(false)
	maps.Execute(false)
	methods.Execute(false)
	interfaces.Execute(true)
	send_mail.Execute(false)
	// user_register.Register()
	// player.Menu()

	// Examples

	sum.Execute(false)
}