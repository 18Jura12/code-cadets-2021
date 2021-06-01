package controllers

import (
	"context"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
)

type BetService interface {
	GetBetById(ctx context.Context ,id string) (domainmodels.BetResponseDto, error)
	GetBetsByUserId(ctx context.Context ,userId string) ([]domainmodels.BetResponseDto, error)
	GetBetsByStatus(ctx context.Context ,status string) ([]domainmodels.BetResponseDto, error)
}
