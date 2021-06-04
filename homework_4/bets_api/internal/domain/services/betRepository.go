package services

import (
	"context"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
)

type BetRepository interface {
	GetBetById(ctx context.Context, id string) (domainmodels.BetResponseDto, bool, error)
	GetBetsByUserId(ctx context.Context, userId string) ([]domainmodels.BetResponseDto, bool, error)
	GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.BetResponseDto, bool, error)
}
