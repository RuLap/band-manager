package repository

import (
	"band-manager/services/user-service/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (uuid.UUID, error) {
	query := `INSERT INTO users (first_name, last_name, email, password, photo_url) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query,
		user.FirstName, user.LastName, user.Email, user.Password, user.PhotoUrl).Scan(&id)
	return id, err
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.PhotoUrl)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	row := r.db.QueryRow(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.PhotoUrl)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
