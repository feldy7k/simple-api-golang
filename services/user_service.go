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

func AddUser(user models.User) {
	
}

func RemoveUser(id string) {
	
}