package controllers

import "github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/models"

type BetValidator interface {
	IsBetValid(betRequest models.BetRequestDto) (bool, error)
}
