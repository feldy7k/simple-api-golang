package services

import (
	"context"
	"go-api/config"
	"go-api/models"
)

func GetUsers() ([]models.User, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err  // <-- return two values here: nil slice + error
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err  // <-- again two values
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id string) (*models.User, error) {
	var user models.User

	err := config.DB.QueryRow(context.Background(),
		"SELECT id, name, email FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err // other errors
	}

	return &user, nil
}

func CreateUser(user models.User) error {
	query := "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)"
	_, err := config.DB.Exec(context.Background(), query, user.ID, user.Name, user.Email)
	return err
}

func UpdateUser(user models.User) error {
	query := "UPDATE users SET name=$2, email=$3 WHERE id=$1"
	_, err := config.DB.Exec(context.Background(), query, user.ID, user.Name, user.Email)
	return err
}

func RemoveUser(id string) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := config.DB.Exec(context.Background(), query, id)
	return err
}