package band

import (
	"band-manager/services/band-service/internal/dto"
	"band-manager/services/band-service/internal/models"
)

func CreateDTOToEntity(dto *dto.CreateBandDTO) models.Band {
	return models.Band{
		Name:     dto.Name,
		PhotoUrl: dto.PhotoURL,
	}
}

func EntityToInfoDTO(band *models.Band, members []dto.MemberInfoDTO) dto.BandInfoDTO {
	return dto.BandInfoDTO{
		ID:       band.ID,
		Name:     band.Name,
		PhotoURL: band.PhotoUrl,
		Members:  members,
	}
}
