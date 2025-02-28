package dto

import "github.com/google/uuid"

type CreateBandDTO struct {
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url"`
}

type BandInfoDTO struct {
	ID       uuid.UUID       `json:"id"`
	Name     string          `json:"name"`
	PhotoURL string          `json:"photo_url"`
	Members  []MemberInfoDTO `json:"members"`
}
