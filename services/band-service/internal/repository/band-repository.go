package repository

import (
	"band-manager/services/band-service/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BandRepository interface {
	Create(ctx context.Context, band *models.Band) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Band, error)
}

type bandRepository struct {
	db *pgxpool.Pool
}

func NewBandRepository(db *pgxpool.Pool) BandRepository {
	return &bandRepository{db: db}
}

func (r *bandRepository) Create(ctx context.Context, band *models.Band) (uuid.UUID, error) {
	query := `INSERT INTO bands (name, photo_url) 
              VALUES ($1, $2) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, band.Name, band.PhotoUrl).Scan(&id)
	return id, err
}

func (r *bandRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Band, error) {
	query := `SELECT * FROM bands WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var band models.Band
	err := row.Scan(&band.ID, &band.Name, &band.PhotoUrl)
	if err != nil {
		return nil, errors.New("band not found")
	}
	return &band, nil
}
