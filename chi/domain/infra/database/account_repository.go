package database

import "venzel.com.br/chi/domain/account"

type AccountRepository struct{}

func (a *AccountRepository) Create(account *account.Account) error {
	return nil
}
