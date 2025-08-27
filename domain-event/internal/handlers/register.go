package handlers

import (
	"context"
	"domain-event/internal/am"
	"domain-event/internal/ddd"
	"domain-event/internal/domain"
)

func RegisterProducedEventHandlers(eventDispatcher *ddd.EventDispatcher, producedEventHandlers *ProducedEventHandlers[ddd.Event]) {
	eventDispatcher.Subcribe(domain.StoreCreatedEvent, producedEventHandlers.HandleEvent)
}

func RegisterConsumerEventHandler(stream am.EventSubscriber, storeHandlers ddd.EventHandler) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return storeHandlers(ctx, eventMsg)
	})

	return stream.Subscribe(domain.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		domain.StoreCreatedEventAsync,
	}, am.GroupName("baskets-stores"))
}
