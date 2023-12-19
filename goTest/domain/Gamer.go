package domain

type Hobbies struct {
	List []string
}

type Gamer struct {
	Name    string
	Age     uint8
	Hobbies *Hobbies
}
