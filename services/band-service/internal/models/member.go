package models

import (
	"github.com/google/uuid"
	"time"
)

type Member struct {
	ID        uuid.UUID `json:"id"`
	BandID    uuid.UUID `json:"band_id"`
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"add_date"`
}
