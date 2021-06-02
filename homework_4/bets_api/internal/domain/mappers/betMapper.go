package mappers

import (
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite/models"
)

type BetMapper struct{}

func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

func (m *BetMapper) MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.BetResponseDto {
	return domainmodels.BetResponseDto{
		Id:                   storageBet.Id,
		Status:               storageBet.Status,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
		Payout:               float64(storageBet.Payout) / 100,
	}
}
