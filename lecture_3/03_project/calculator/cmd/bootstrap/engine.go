package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine/consumer"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine/handler"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine/publisher"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite"
)

func newBetConsumer(channel rabbitmq.Channel) *rabbitmq.BetConsumer {
	betConsumer, err := rabbitmq.NewBetConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerBetQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerBetName,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		},
	)
	if err != nil {
		panic(err)
	}
	return betConsumer
}

func newEventUpdateConsumer(channel rabbitmq.Channel) *rabbitmq.EventUpdateConsumer {
	eventUpdateConsumer, err := rabbitmq.NewEventUpdateConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerEventUpdateQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerEventUpdateName,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		},
	)
	if err != nil {
		panic(err)
	}
	return eventUpdateConsumer
}

func newConsumer(betConsumer consumer.BetConsumer, eventUpdateConsumer consumer.EventUpdateConsumer) *consumer.Consumer {
	return consumer.New(betConsumer, eventUpdateConsumer)
}

func newCalcBetMapper() *mappers.CalcBetMapper {
	return mappers.NewCalcBetMapper()
}

func newCalcBetRepository(dbExecutor sqlite.DatabaseExecutor, calcBetMapper sqlite.CalcBetMapper) *sqlite.CalcBetRepository {
	return sqlite.NewCalcBetRepository(dbExecutor, calcBetMapper)
}

func newHandler(calcBetRepository handler.CalcBetRepository) *handler.Handler {
	return handler.New(calcBetRepository)
}

func newBetCalculatedPublisher(channel rabbitmq.Channel) *rabbitmq.BetCalculatedPublisher {
	betCalculatedPublisher, err := rabbitmq.NewBetCalculatedPublisher(
		channel,
		rabbitmq.PublisherConfig{
			Queue:             config.Cfg.Rabbit.PublisherBetCalculatedQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			Exchange:          config.Cfg.Rabbit.PublisherExchange,
			Mandatory:         config.Cfg.Rabbit.PublisherMandatory,
			Immediate:         config.Cfg.Rabbit.PublisherImmediate,
		},
	)
	if err != nil {
		panic(err)
	}
	return betCalculatedPublisher
}

func newPublisher(betCalculatedPublisher publisher.BetCalculatedPublisher) *publisher.Publisher {
	return publisher.New(betCalculatedPublisher)
}

func Engine(rabbitMqChannel rabbitmq.Channel, dbExecutor sqlite.DatabaseExecutor) *engine.Engine {
	betConsumer := newBetConsumer(rabbitMqChannel)
	eventUpdateConsumer := newEventUpdateConsumer(rabbitMqChannel)
	consumer := newConsumer(betConsumer, eventUpdateConsumer)

	calcBetMapper := newCalcBetMapper()
	calcBetRepository := newCalcBetRepository(dbExecutor, calcBetMapper)
	handler := newHandler(calcBetRepository)

	betPublisher := newBetCalculatedPublisher(rabbitMqChannel)
	publisher := newPublisher(betPublisher)

	return engine.New(consumer, handler, publisher)
}

