package dto

import "github.com/google/uuid"

type AddMemberDTO struct {
	BandID uuid.UUID `json:"band_id"`
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}

type MemberInfoDTO struct {
	ID   uuid.UUID   `json:"id"`
	Band BandInfoDTO `json:"band"`
	User UserInfoDTO `json:"user"`
	Role string      `json:"role"`
}

type UserInfoDTO struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	PhotoUrl  string    `json:"photo_url"`
}
