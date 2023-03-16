package main

import (
	"fmt"
	"inter/users"
)

func main() {
	professor := users.Professor{Name: "tiago"}
	// aluno := users.Aluno{Name: "marcos"}

	createProfessor := users.CreateUser(&professor)
	// createAluno := users.CreateUser(&aluno)

	fmt.Println(createProfessor)
	// fmt.Println(createAluno)
}
