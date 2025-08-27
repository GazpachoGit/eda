package ddd

import (
	"context"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.Mutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher) Subcribe(eventName string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventName] = append(d.handlers[eventName], handler)
}

func (d *EventDispatcher) Publish(ctx context.Context, event Event) error {
	name := event.EventName()
	handlers, ok := d.handlers[name]
	if ok {
		for _, h := range handlers {
			if err := h(ctx, event); err != nil {
				return err
			}
		}
	}
	return nil
}
