package services

import (
	"band-manager/services/band-service/internal/dto"
	band_mapper "band-manager/services/band-service/internal/mapper/custom/band"
	member_mapper "band-manager/services/band-service/internal/mapper/custom/member"
	"band-manager/services/band-service/internal/repository"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type BandService interface {
	Create(ctx context.Context, bandDto *dto.CreateBandDTO) (*dto.BandInfoDTO, error)
	GetByID(ctx context.Context, id string) (*dto.BandInfoDTO, error)
	AddMember(ctx context.Context, memberDto *dto.AddMemberDTO) (*dto.BandInfoDTO, error)
}

type bandService struct {
	bandRepo   repository.BandRepository
	memberRepo repository.MemberRepository
}

func NewBandService(bandRepo repository.BandRepository, memberRepo repository.MemberRepository) BandService {
	return &bandService{bandRepo: bandRepo, memberRepo: memberRepo}
}

func (s *bandService) Create(ctx context.Context, bandDto *dto.CreateBandDTO) (*dto.BandInfoDTO, error) {
	band := band_mapper.CreateDTOToEntity(bandDto)
	band.ID = uuid.New()

	createdId, err := s.bandRepo.Create(ctx, &band)
	if err != nil {
		slog.Error("failed to create band", "error", err)
		return nil, err
	}

	created, err := s.bandRepo.GetByID(ctx, createdId)
	if err != nil {
		slog.Error("failed to get created band by id", "error", err)
		return nil, err
	}

	//TODO: Add Members loading
	cratedDto := band_mapper.EntityToInfoDTO(created, nil)

	return &cratedDto, nil
}

func (s *bandService) GetByID(ctx context.Context, id string) (*dto.BandInfoDTO, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	band, err := s.bandRepo.GetByID(ctx, uid)
	if err != nil {
		slog.Error("failed to get band", "bandID", id, "error", err)
	}

	//TODO: Add Members loading
	bandDto := band_mapper.EntityToInfoDTO(band, nil)

	return &bandDto, nil
}

func (s *bandService) AddMember(ctx context.Context, memberDto *dto.AddMemberDTO) (*dto.BandInfoDTO, error) {
	member := member_mapper.CreateDTOToEntity(memberDto)
	_, err := s.memberRepo.Create(ctx, &member)
	if err != nil {
		slog.Error("failed to create member", "error", err)
		return nil, err
	}

	band, err := s.bandRepo.GetByID(ctx, memberDto.BandID)
	if err != nil {
		slog.Error("failed to create member", "error", err)
		return nil, err
	}

	//TODO: Add Members loading
	bandDto := band_mapper.EntityToInfoDTO(band, nil)

	return &bandDto, nil
}
