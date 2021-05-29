package controllers

import "github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite/models"

type BetService interface {
	GetBetById(id string) (models.Bet, error)
	GetBetsByUserId() ([]models.Bet, error)
	GetBetsByStatus() ([]models.Bet, error)
}
