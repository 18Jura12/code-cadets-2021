package services

import (
	"context"
	"github.com/pkg/errors"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
)

type BetService struct {
	betRepository BetRepository
}

func NewBetService(betRepository BetRepository) *BetService {
	return &BetService{
		betRepository: betRepository,
	}
}

func (s *BetService) GetBetById(ctx context.Context ,id string) (domainmodels.BetResponseDto, error) {
	resultingBet, exists, err := s.betRepository.GetBetById(ctx, id)
	if !exists {
		err = errors.WithMessage(err, "no such bet")
	}
	return resultingBet, err
}

func (s *BetService) GetBetsByUserId(ctx context.Context, userId string) ([]domainmodels.BetResponseDto, error) {
	resultingBets, exists, err := s.betRepository.GetBetsByUserId(ctx, userId)
	if !exists {
		err = errors.WithMessage(err, "no such bets")
	}
	return resultingBets, err
}

func (s *BetService) GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.BetResponseDto, error) {
	resultingBets, exists, err := s.betRepository.GetBetsByStatus(ctx, status)
	if !exists {
		err = errors.WithMessage(err, "no such bets")
	}
	return resultingBets, err
}
