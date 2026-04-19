package service

import (
	"errors"
	"fmt"

	"github.com/aigerim25/Goo/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) GetUserById(id int) (*repository.User, error) {
	return s.repo.GetUserById(id)
}
func (s *UserService) CreateUser(user *repository.User) error {
	return s.repo.CreateUser(user)
}
func (s *UserService) RegisterUser(user *repository.User, email string) error {
	existing, err := s.repo.GetByEmail(email)
	if existing != nil {
		return fmt.Errorf("user with this email already exists")
	}
	if err != nil {
		return fmt.Errorf("error getting user with this email")
	}
	return s.repo.CreateUser(user)
}
func (s *UserService) UpdateUserName(id int, newName string) error {
	if newName == "" {
		return fmt.Errorf("new name cannot be empty")
	}
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return err
	}
	user.Name = newName
	return s.repo.UpdateUser(user)
}
func (s *UserService) DeleteUser(id int) error {
	if id == 1 {
		return errors.New("it is not allowed to delete admin user")
	}
	return s.repo.DeleteUser(id)
}
