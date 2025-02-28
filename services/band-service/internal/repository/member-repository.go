package repository

import (
	"band-manager/services/band-service/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MemberRepository interface {
	Create(ctx context.Context, member *models.Member) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Member, error)
}

type memberRepository struct {
	db *pgxpool.Pool
}

func NewMemberRepository(db *pgxpool.Pool) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) Create(ctx context.Context, member *models.Member) (uuid.UUID, error) {
	query := `INSERT INTO members (band_id, user_id, role) 
              VALUES ($1, $2, $3) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, member.BandID, member.UserID, member.Role).Scan(&id)
	return id, err
}

func (r *memberRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Member, error) {
	query := `SELECT * FROM members WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var member models.Member
	err := row.Scan(&member.ID, &member.BandID, &member.UserID, &member.Role)
	if err != nil {
		return nil, errors.New("band member not found")
	}
	return &member, nil
}
