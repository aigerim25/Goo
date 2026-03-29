package repository

import (
	"database/sql"
	"fmt"
	"practice5/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetPaginatedUsers(page, pageSize int, filters map[string]string, orderBy string) (models.PaginatedResponse, error) {

	var users []models.User
	offset := (page - 1) * pageSize

	query := "SELECT id, name, email, gender, birth_date FROM users WHERE 1 = 1"
	args := []interface{}{}
	argIndex := 1

	if v, ok := filters["id"]; ok {
		query += fmt.Sprintf(" AND id = $%d", argIndex)
		args = append(args, v)
		argIndex++
	}
	if v, ok := filters["name"]; ok {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, "%"+v+"%")
		argIndex++
	}
	if v, ok := filters["email"]; ok {
		query += fmt.Sprintf(" AND email ILIKE $%d", argIndex)
		args = append(args, "%"+v+"%")
		argIndex++
	}
	if v, ok := filters["gender"]; ok {
		query += fmt.Sprintf(" AND gender = $%d", argIndex)
		args = append(args, v)
		argIndex++
	}
	if v, ok := filters["birth_date"]; ok {
		query += fmt.Sprintf(" AND birth_date = $%d", argIndex)
		args = append(args, v)
		argIndex++
	}
	if orderBy != "" {
		query += " ORDER BY " + orderBy
	} else {
		query += " ORDER BY id "
	}
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pageSize, offset)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return models.PaginatedResponse{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Gender,
			&u.BirthDate,
		)
		if err != nil {
			return models.PaginatedResponse{}, err
		}
		users = append(users, u)
	}
	return models.PaginatedResponse{
		Data:     users,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
func (r *UserRepository) GetCommonFriends(userID1 int, userID2 int) ([]models.User, error) {
	friends := make([]models.User, 0)
	query := `SELECT u.id, u.name, u.email, u.gender, u.birth_date FROM user_friends uf1
	          JOIN user_friends uf2 ON uf1.friend_id = uf2.friend_id
			  JOIN users u ON u.id = uf1.friend_id
			  WHERE uf1.user_id = $1 AND uf2.user_id = $2`
	rows, err := r.db.Query(query, userID1, userID2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate)
		if err != nil {
			return nil, err
		}
		friends = append(friends, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return friends, nil
}
