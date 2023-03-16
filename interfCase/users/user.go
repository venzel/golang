package users

type User interface {
	Create() User
}

func CreateUser(u User) User {

	return u.Create()
}
