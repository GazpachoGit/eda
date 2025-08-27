package handlers

import (
	"context"
	"domain-event/internal/am"
	"domain-event/internal/ddd"
	"domain-event/internal/domain"
)

type ProducedEventHandlers[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

func NewProducedEventHandlers(publisher am.MessagePublisher[ddd.Event]) *ProducedEventHandlers[ddd.Event] {
	return &ProducedEventHandlers[ddd.Event]{
		publisher: publisher,
	}
}

func (h *ProducedEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case domain.StoreParticipatingToggledEvent:
		return h.onStoreParticipationChange(ctx, event)
	}
	return nil
}

func (h ProducedEventHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreCreated)
	return h.publisher.Publish(ctx, domain.StoreAggregateChannel,
		ddd.NewEvent(domain.StoreCreatedEventAsync, payload),
	)
}

func (h ProducedEventHandlers[T]) onStoreParticipationChange(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreParticipationToggled)
	return h.publisher.Publish(ctx, domain.StoreAggregateChannel,
		ddd.NewEvent(domain.StoreParticipatingToggledEventAsync, payload),
	)
}
