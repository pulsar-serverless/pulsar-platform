package analytics

import (
	"context"
	"encoding/json"
	domain "pulsar/internal/core/domain/analytics"
)

func (service *analyticsService) PublishInvocationCreatedEvent(ctx context.Context, invocation *domain.Invocation) error {
	return service.queue.Publish(ctx, service.createInvocationQueue, invocation)
}

func (service *analyticsService) consumeInvocationCreatedEvent(ctx context.Context) {
	service.queue.Consume(context.Background(), service.createInvocationQueue, func(message []byte) error {
		var log domain.Invocation

		err := json.Unmarshal(message, &log)
		if err != nil {
			return err
		}

		return service.invocationRepo.CreateInvocation(ctx, &log)
	})
}
