package user

import (
	"band-manager/services/user-service/internal/dto"
	"band-manager/services/user-service/internal/models"
)

func RegisterDTOToEntity(dto *dto.UserRegisterDTO) models.User {
	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}
}

func EntityToInfoDTO(user *models.User) dto.UserInfoDTO {
	return dto.UserInfoDTO{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		PhotoUrl:  user.PhotoUrl,
	}
}

func UserToLoginDTO(user *dto.UserInfoDTO, token string) dto.UserLoginResponseDTO {
	return dto.UserLoginResponseDTO{
		User:  *user,
		Token: token,
	}
}
