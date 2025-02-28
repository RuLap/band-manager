package services

import (
	"band-manager/services/band-service/internal/dto"
	member_mapper "band-manager/services/band-service/internal/mapper/custom/member"
	"band-manager/services/band-service/internal/repository"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type MemberService interface {
	Create(ctx context.Context, memberDto *dto.AddMemberDTO) (*dto.MemberInfoDTO, error)
	GetByID(ctx context.Context, id string) (*dto.MemberInfoDTO, error)
}

type memberService struct {
	memberRepo repository.MemberRepository
}

func NewMemberService(memberRepo repository.MemberRepository) MemberService {
	return &memberService{memberRepo: memberRepo}
}

func (s *memberService) Create(ctx context.Context, memberDto *dto.AddMemberDTO) (*dto.MemberInfoDTO, error) {
	member := member_mapper.CreateDTOToEntity(memberDto)
	member.ID = uuid.New()

	createdId, err := s.memberRepo.Create(ctx, &member)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return nil, err
	}

	created, err := s.memberRepo.GetByID(ctx, createdId)
	if err != nil {
		slog.Error("failed to get created member by id", "error", err)
		return nil, err
	}

	//TODO: Add User loading from user service
	cratedDto := member_mapper.EntityToInfoDTO(created, nil, nil)

	return &cratedDto, nil
}

func (s *memberService) GetByID(ctx context.Context, id string) (*dto.MemberInfoDTO, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	member, err := s.memberRepo.GetByID(ctx, uid)
	if err != nil {
		slog.Error("failed to get user", "userID", id, "error", err)
	}

	//TODO: Add User loading from user service
	userDto := member_mapper.EntityToInfoDTO(member, nil, nil)

	return &userDto, nil
}
