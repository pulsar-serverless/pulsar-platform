package ports

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IMessageQueue interface {
	Publish(ctx context.Context, queue string, data interface{}) error
	Consume(ctx context.Context, queue string, handler func(message []byte) error)
	CreateQueue(queueName string) amqp.Queue
}
