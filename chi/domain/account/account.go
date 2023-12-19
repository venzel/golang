package account

type Account struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Account) NewAccount(name, email, password string) (*Account, error) {
	return &Account{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
