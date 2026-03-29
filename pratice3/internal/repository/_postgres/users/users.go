package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"practice3/internal/repository/_postgres"
	"practice3/pkg/modules"
)

var ErrUserNotFound = errors.New("user not found")

type Repository struct {
	db               *_postgres.Dialect
	executionTimeout time.Duration
}

func NewUserRepository(db *_postgres.Dialect) *Repository {
	return &Repository{
		db:               db,
		executionTimeout: time.Second * 5,
	}
}
func (r *Repository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.DB.Select(&users, "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}
func (r *Repository) CreateUser(u modules.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.executionTimeout)
	defer cancel()

	if u.Name == "" {
		return 0, fmt.Errorf("user name is required")
	}
	if u.Email == "" {
		return 0, fmt.Errorf("user email is required")
	}
	if u.Phone == "" {
		return 0, fmt.Errorf("user phone is required")
	}
	if u.Age < 0 {
		return 0, fmt.Errorf("age must be greater than zero")
	}
	var id int
	err := r.db.DB.QueryRowContext(
		ctx,
		`INSERT INTO users(name, email, age, phone) VALUES ($1, $2, $3, $4)
         RETURNING id`,
		u.Name, u.Email, u.Age, u.Phone).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *Repository) UpdateUser(u modules.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.executionTimeout)
	defer cancel()

	if u.ID <= 0 {
		return fmt.Errorf("user id must be greater than zero")
	}
	if u.Name == "" {
		return fmt.Errorf("user name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("user email is required")
	}
	if u.Phone == "" {
		return fmt.Errorf("user phone is required")
	}
	if u.Age < 0 {
		return fmt.Errorf("age must be greater than zero")
	}
	res, err := r.db.DB.ExecContext(
		ctx,
		`UPDATE users
         SET name = $1, email = $2, age = $3, phone = $4
         WHERE id = $5`,
		u.Name, u.Email, u.Age, u.Phone, u.ID)
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil

}
