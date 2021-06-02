package services

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/models"
)

type BetService struct {
	betPublisher BetPublisher
}

func NewBetService(betPublisher BetPublisher) *BetService {
	return &BetService{
		betPublisher: betPublisher,
	}
}

func (b *BetService) CreateBet(bet models.BetRequestDto) error {
	return b.betPublisher.Publish(bet)
}
