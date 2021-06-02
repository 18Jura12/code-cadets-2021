package services

import "github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/models"

type BetPublisher interface {
	Publish(bet models.BetRequestDto) error
}

