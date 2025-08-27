package handlers

import (
	"context"
	"domain-event/internal/am"
	"domain-event/internal/ddd"
	"domain-event/internal/domain"
)

type IntegrationEventHandlers[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.Event] {
	return &IntegrationEventHandlers[ddd.Event]{
		publisher: publisher,
	}
}

func (h *IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case domain.StoreParticipatingToggledEvent:
		return h.onStoreParticipationChange(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreCreated)
	return h.publisher.Publish(ctx, domain.StoreAggregateChannel,
		ddd.NewEvent(domain.StoreCreatedEventAsync, payload),
	)
}

func (h IntegrationEventHandlers[T]) onStoreParticipationChange(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreParticipationToggled)
	return h.publisher.Publish(ctx, domain.StoreAggregateChannel,
		ddd.NewEvent(domain.StoreParticipatingToggledEventAsync, payload),
	)
}
