package rabbitmq

import "github.com/streadway/amqp"

type QueuePublisher interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}

