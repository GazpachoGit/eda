package handlers

import (
	"context"
	"domain-event/internal/ddd"
	"domain-event/internal/domain"
	"fmt"
)

type ConsumedEventHandler struct {
}

func (h ConsumedEventHandler) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	}
	return nil
}

func (h ConsumedEventHandler) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.StoreCreated)
	fmt.Println("I recieved a 'store created' event : ", payload.Name)
	return nil
}

func NewConsumedEventHandler() ConsumedEventHandler {
	return ConsumedEventHandler{}
}
