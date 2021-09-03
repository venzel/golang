package player

import "fmt"

func choice_add(p *[]gamer) {
	var name string
	var wins int
	var other string

	for {
		fmt.Print("Digite seu nome: ")
		fmt.Scan(&name)
		
		fmt.Print("Digite o total de wins: ")
		fmt.Scan(&wins)

		add(name, wins, p)

		fmt.Print("Deseja adicionar outro? [s|n] ")
		fmt.Scan(&other)

		if other != "s" {
			break
		}
	}
}

func Menu() {
	list := []gamer{}

	choice_add(&list)

	print(&list)
}