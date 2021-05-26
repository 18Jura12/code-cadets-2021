package publisher

import (
	"context"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Publisher struct {
	betCalculatedPublisher BetCalculatedPublisher
}

func New(betPublisher BetCalculatedPublisher) *Publisher {
	return &Publisher{
		betCalculatedPublisher: betPublisher,
	}
}

func (p *Publisher) PublishBetsCalculated(ctx context.Context, bets <-chan rabbitmqmodels.BetCalculated) {
	p.betCalculatedPublisher.Publish(ctx, bets)
}