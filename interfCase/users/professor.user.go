package users

type Professor struct {
	Name string
}

func (p *Professor) Create() User {
	return &Professor{
		Name: "tiago",
	}
}
