package publisher

import (
	"context"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type BetCalculatedPublisher interface {
	Publish(ctx context.Context, betsCalculated <-chan models.BetCalculated)
}
