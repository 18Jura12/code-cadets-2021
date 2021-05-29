package services

import (
	"github.com/gin-gonic/gin"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite/models"
)

type BetService struct {
	betRepository BetRepository
}

func NewBetService(betRepository BetRepository) *BetService {
	return &BetService{
		betRepository: betRepository,
	}
}

func (s *BetService) GetBetById(ctx gin.Context ,id string) (domainmodels.BetResponseDto, error) {
	resultingBet, exists, err := s.betRepository.GetBetById(ctx, id)
	if err != nil {

	}
	if !exists {

	}
	return resultingBet, nil
}

func (s *BetService) GetBetsByUserId() ([]models.Bet, error) {

}

func (s *BetService) GetBetsByStatus() ([]models.Bet, error) {

}