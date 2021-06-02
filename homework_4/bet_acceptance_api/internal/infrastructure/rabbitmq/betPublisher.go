package rabbitmq

import (
	"encoding/json"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	models2 "github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/models"
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

func (b *BetPublisher) Publish(betRequest models2.BetRequestDto) error {
	id, err := getRandomUUID()
	if err != nil {
		return err
	}

	bet := &models.BetDto{
		Id:                   id,
		CustomerId:           betRequest.CustomerId,
		SelectionId:          betRequest.SelectionId,
		SelectionCoefficient: betRequest.SelectionCoefficient,
		Payment:              betRequest.Payment,
	}

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

func getRandomUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.WithMessage(err, "failed to generate a uuid")
	}
	return id.String(), nil
}

