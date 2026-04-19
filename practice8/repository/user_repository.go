package repository

type User struct {
	ID   int
	Name string
}
type UserRepository interface {
	GetUserById(id int) (*User, error)
	CreateUser(user *User) error
	GetByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

// на заметку команду mockgen -source ... пишем без =
