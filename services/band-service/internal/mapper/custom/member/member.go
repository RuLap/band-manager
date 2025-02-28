package member

import (
	"band-manager/services/band-service/internal/dto"
	"band-manager/services/band-service/internal/models"
)

func CreateDTOToEntity(dto *dto.AddMemberDTO) models.Member {
	return models.Member{
		BandID: dto.BandID,
		UserID: dto.UserID,
		Role:   dto.Role,
	}
}

func EntityToInfoDTO(member *models.Member, band *dto.BandInfoDTO, user *dto.UserInfoDTO) dto.MemberInfoDTO {
	return dto.MemberInfoDTO{
		ID:   member.ID,
		Band: *band,
		User: *user,
		Role: member.Role,
	}
}
