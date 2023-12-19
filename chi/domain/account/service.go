package account

type Service interface {
	CreateAccount() error
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) CreateAccount() error {
	account := Account{}

	obj, err := account.NewAccount("name", "email", "password")

	if err != nil {
		panic("erro")
	}

	return s.Repository.Create(obj)
}
