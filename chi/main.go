package main

import (
	"fmt"

	"venzel.com.br/chi/domain/account"
	"venzel.com.br/chi/domain/infra/database"
)

func main() {

	accountService := account.ServiceImpl{
		Repository: &database.AccountRepository{},
	}

	fmt.Println(accountService)
}
