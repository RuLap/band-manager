package models

import "github.com/google/uuid"

type Band struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	PhotoUrl string    `json:"photoUrl"`
}
