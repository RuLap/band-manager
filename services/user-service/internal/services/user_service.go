package services

import (
	"band-manager/services/user-service/internal/dto"
	user_mapper "band-manager/services/user-service/internal/mapper/custom/user"
	"band-manager/services/user-service/internal/repository"
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrEmailExists     = errors.New("email exists")
	ErrInvalidPassword = errors.New("invalid password")
)

type UserService interface {
	Register(ctx context.Context, user *dto.UserRegisterDTO) (*dto.UserInfoDTO, error)
	Login(ctx context.Context, email string, password string) (*dto.UserInfoDTO, error)
	GetByID(ctx context.Context, id string) (*dto.UserInfoDTO, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserInfoDTO, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, userDto *dto.UserRegisterDTO) (*dto.UserInfoDTO, error) {
	existingUser, _ := s.userRepo.GetByEmail(ctx, userDto.Email)
	if existingUser != nil {
		slog.Error("failed to register with existing email", "error", ErrEmailExists)
		return nil, ErrEmailExists
	}

	user := user_mapper.RegisterDTOToEntity(userDto)
	user.ID = uuid.NewString()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to generate password hash", "error", err)
		return nil, err
	}
	user.Password = string(hashedPassword)

	createdId, err := s.userRepo.Create(ctx, &user)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return nil, err
	}

	created, err := s.userRepo.GetByID(ctx, createdId)
	if err != nil {
		slog.Error("failed to get created user by id", "error", err)
		return nil, err
	}

	cratedDto := user_mapper.EntityToInfoDTO(created)

	return &cratedDto, nil
}

func (s *userService) Login(ctx context.Context, email string, password string) (*dto.UserInfoDTO, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("failed to find user", "error", ErrUserNotFound)
		return nil, ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		slog.Error("failed to compare passwords", "error", ErrInvalidPassword)
		return nil, ErrInvalidPassword
	}

	userDto := user_mapper.EntityToInfoDTO(user)

	return &userDto, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*dto.UserInfoDTO, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		slog.Error("failed to get user", "userID", id, "error", err)
	}

	userDto := user_mapper.EntityToInfoDTO(user)

	return &userDto, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*dto.UserInfoDTO, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("failed to find user", "email", email, "error", ErrEmailExists)
		return nil, err
	}

	userDto := user_mapper.EntityToInfoDTO(user)

	return &userDto, nil
}
