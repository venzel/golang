package main

import (
	"log"
	"module/examples/cli/app"
	"os"
)

func main() {
	application := app.Generate()

	/*
	erro := application.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
	*/

	// infinity, a mesma coisa que o codigo acima
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}