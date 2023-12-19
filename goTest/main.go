package main

func main() {
	dbz := make(map[string]uint8)

	dbz["Goku"] = 40
	dbz["Vegeta"] = 42

	_, exists := dbz["Goku"]

	if exists {
		println("Goku exists")
	}

	for name, age := range dbz {
		println(name, age)
	}

	persons := make([]string, 10)

	persons = append(persons, "Goku")
	persons = append(persons, "Vegita")

	for index, person := range persons {
		println(index, person)
	}
}
