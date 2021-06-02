package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/infrastructure/rabbitmq/models"
	"log"
)

const contentTypeTextPlain = "text/plain"

type BetPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

func NewBetPublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetPublisher {
	return &BetPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

func (b *BetPublisher) Publish(bet models.BetDto) error {
	betUpdateJson, err := json.Marshal(bet)
	if err != nil {
		return err
	}

	err = b.publisher.Publish(
		b.exchange,
		b.queueName,
		b.mandatory,
		b.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        betUpdateJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", betUpdateJson)
	return nil
}
