package repository

import (
	"practice3/internal/repository/_postgres"
	"practice3/internal/repository/_postgres/users"
	"practice3/pkg/modules"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
	CreateUser(u modules.User) (int, error)
	UpdateUser(u modules.User) error
}
type Repositories struct {
	UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{
		users.NewUserRepository(db),
	}
}
