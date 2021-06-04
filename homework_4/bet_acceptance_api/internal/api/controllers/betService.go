package controllers

import "github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/domain/models"

type BetService interface {
	CreateBet(bet models.BetRequest) error
}

