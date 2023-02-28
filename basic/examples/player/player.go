package player

import (
	"fmt"
)

func add(name string, wins int, p *[]gamer) {
	(*p) = append((*p), gamer{name,  wins})
}

func size(p *[]gamer) int {
	return len((*p))
}

func empty(p *[]gamer) bool {
	return size(p) < 0
}

func top(p *[]gamer) []gamer {
	list := []gamer{}

	for i := range (*p) {
		if (*p)[i].wins >= 2 {
			list = append(list, (*p)[i])
		}
	}

	return list
}

func position(name string, p *[]gamer) int {
	index := -1

	if empty(p) {
		return index
	}

	for k := range (*p) {
		if (*p)[k].name == name {
			index = k;
			return index
		}
	}

	return index
}

func model(position int, p *[]gamer) gamer {
	return (*p)[position]
}

func print(p *[]gamer) {
	for i := range (*p) {
		fmt.Println((*p)[i].name, (*p)[i].wins)
	}
}

// func Execute() {
// 	list := reference()

// 	add("tiago", 3, &list)
// 	add("liz", 2, &list)
// 	add("cintia", 0, &list)

// 	size := size(&list)

// 	fmt.Println(size)

// 	position := position("liz", &list)

// 	fmt.Println(position)

// 	print(&list)

// 	if position != -1 {
// 		fmt.Println(model(position, &list))
// 	}

// 	top := top(&list)

// 	print(&top)
// }