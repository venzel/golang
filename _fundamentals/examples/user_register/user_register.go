package user_register

import (
	"fmt"
)

type user struct {
	name string
	victories uint8
}

func add(name string, p *[]user) {
	(*p) = append((*p), user{name: name})
}

func exists(p *[]user) bool {
	return len((*p)) > 0
}

func print(p *[]user) {
	for i := range *p {
		fmt.Println((*p)[i].name)
	}
}

func getUserByIndex(index int, p *[]user) user {
	return (*p)[index]
}

func findUserByName(name string, p *[]user) int {
	index := -1

	if !exists(p) {
		return index
	}

	for i := range *p {
		if ((*p)[i].name == name) {
			index = i
			break
		}
	}

	return index
}

func Register() {

	users := []user{
		{"Tiago", 0},
		{"Cintia", 0},
	}

	add("Liz", &users)

	userIndex := findUserByName("Tiago", &users)

	if userIndex == -1 {
		fmt.Println("error!")
	} else {
		user := getUserByIndex(userIndex, &users)
		fmt.Println(user)
	}

	print(&users)
}