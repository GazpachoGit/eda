package handlers

import (
	"domain-event/internal/ddd"
	"domain-event/internal/domain"
)

func RegisterIntegrationEventHandlers(eventDispatcher *ddd.EventDispatcher, integrationEventHandlers *IntegrationEventHandlers[ddd.Event]) {
	eventDispatcher.Subcribe(domain.StoreCreatedEvent, integrationEventHandlers.HandleEvent)
}
