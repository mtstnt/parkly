package services

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetAllUsers() {}
