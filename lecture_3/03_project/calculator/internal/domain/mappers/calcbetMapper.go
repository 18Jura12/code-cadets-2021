package mappers

import (
	"math"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

type CalcBetMapper struct {
}

func NewCalcBetMapper() *CalcBetMapper {
	return &CalcBetMapper{}
}

func (m *CalcBetMapper) MapDomainBetToStorageBet(domainBet domainmodels.Bet) storagemodels.CalcBet {
	return storagemodels.CalcBet{
		Id:                   domainBet.Id,
		SelectionId:          domainBet.SelectionId,
		SelectionCoefficient: int(math.Round(domainBet.SelectionCoefficient * 100)),
		Payment:              int(math.Round(domainBet.Payment * 100)),
	}
}

func (m *CalcBetMapper) MapStorageBetToDomainBet(storageBet storagemodels.CalcBet) domainmodels.Bet {
	return domainmodels.Bet{
		Id:                   storageBet.Id,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
	}
}
