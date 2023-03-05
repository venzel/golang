package gamer

import "time"

type Account struct {
	Name string
	Address Address
	Birthday time.Time
}

func (account Account) GetBirthday() int {
	dateInitial := account.Birthday.Year();
	dateFinal := time.Now().Year();
	return dateFinal - dateInitial;
}