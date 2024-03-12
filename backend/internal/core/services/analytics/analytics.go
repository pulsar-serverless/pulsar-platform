package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/ports"
)

type analyticsService struct {
	queue                 ports.IMessageQueue
	createInvocationQueue string
	invocationRepo        ports.InvocationRepository
}

type IAnalyticsService interface {
	PublishInvocationCreatedEvent(ctx context.Context, invocation *domain.Invocation) error
}

func NewAnalyticsService(invocationRepo ports.InvocationRepository, mq ports.IMessageQueue) *analyticsService {
	queue := mq.CreateQueue("CREATE_INVOCATION_QUEUE")

	service := &analyticsService{
		queue:                 mq,
		createInvocationQueue: queue.Name,
		invocationRepo:        invocationRepo,
	}

	go service.consumeInvocationCreatedEvent(context.TODO())

	return service
}
