package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type messageQueue struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewMessageQueue(rabbitMQUrl string) *messageQueue {
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %s", err))
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Failed to open a channel: %s", err))
	}

	return &messageQueue{conn, channel}
}

func (mq *messageQueue) Publish(ctx context.Context, queue string, data interface{}) error {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json error: %s", err)
	}

	return mq.channel.PublishWithContext(ctx, "", queue, true, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        json,
	})
}

func (mq *messageQueue) Consume(ctx context.Context, queue string, handler func(message []byte) error) {
	messages, err := mq.channel.Consume(
		queue, // queue name
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		panic(fmt.Sprintf("Error consuming a channel: %v", err))
	}

	wait := make(chan bool)
	go func() {
		for message := range messages {
			handler(message.Body)
		}

	}()
	<-wait
}

func (mq *messageQueue) CreateQueue(queueName string) amqp.Queue {
	queue, err := mq.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		true,      // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		panic(fmt.Sprintf("Unable to create a queue: %v", err))
	}

	return queue
}
