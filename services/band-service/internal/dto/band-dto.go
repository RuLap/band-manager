package dto

import "github.com/google/uuid"

type CreateBandDTO struct {
	Name string `json:"name"`
}

type BandInfoDTO struct {
	ID      uuid.UUID       `json:"id"`
	Name    string          `json:"name"`
	Members []MemberInfoDTO `json:"members"`
}
