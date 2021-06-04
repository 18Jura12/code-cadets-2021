package bootstrap

import (
	"github.com/streadway/amqp"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/validators"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/infrastructure/rabbitmq"
)

func newBetValidator() *validators.BetValidator {
	return validators.NewBetValidator(
		config.Cfg.ConstVariables.SelectionCoefficientUpperBound,
		config.Cfg.ConstVariables.PaymentLowerBound,
		config.Cfg.ConstVariables.PaymentUpperBound,
	)
}

func newBetPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetPublisher {
	return rabbitmq.NewBetPublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherBetQueueQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newBetService(publisher services.BetPublisher) *services.BetService {
	return services.NewBetService(publisher)
}

func newController(
	betValidator controllers.BetValidator,
	betService controllers.BetService,
) *controllers.Controller {
	return controllers.NewController(betValidator, betService)
}

func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betValidator := newBetValidator()
	betPublisher := newBetPublisher(rabbitMqChannel)
	betService := newBetService(betPublisher)
	controller := newController(betValidator, betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}

