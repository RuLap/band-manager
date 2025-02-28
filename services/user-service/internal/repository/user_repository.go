package repository

import (
	"band-manager/services/user-service/internal/models"
	"context"
)

func CreateUser(user models.User) (int, error) {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(context.Background(), query, user.Email, user.Password).Scan(&id)
	return id, err
}
