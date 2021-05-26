package handler

import (
	"context"
	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
)

type CalcBetRepository interface {
	InsertCalcBet(ctx context.Context, bet domainmodels.Bet) error
	UpdateCalcBet(ctx context.Context, bet domainmodels.Bet) error
	DeleteCalcBet(ctx context.Context, id string) error
	GetCalcBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error)
	GetCalcBetsBySelectionID(ctx context.Context, selectionId string) ([]domainmodels.Bet, bool, error)
}
