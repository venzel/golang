package gamer

import "time"

type Account struct {
	Name string
	Birthday time.Time
	Age int
	Address []Address
}

func (account Account) GetBirthday() (int, time.Time) {
	dateInitial := account.Birthday.Year()
	dateFinal := time.Now().Year()
	return dateFinal - dateInitial, account.Birthday
}

func (account *Account) SetAge(age int) {
	account.Age = age
}
