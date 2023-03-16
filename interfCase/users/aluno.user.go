package users

type Aluno struct {
	Name string
}

func (p *Aluno) Create(name *string) string {
	alu := Aluno{
		Name: *name,
	}

	return alu.Name + " ALUNO!"
}
