package ddd

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Event interface {
	IDer
	EventName() string
	Payload() EventPayload
	Metadata() Metadata
	OccurredAt() time.Time
}

type EventPayload interface{}

type event struct {
	Entity
	payload    EventPayload
	metadata   Metadata
	occurredAt time.Time
}

type EventHandler func(ctx context.Context, event Event) error

func NewEvent(name string, payload EventPayload) event {
	evt := event{
		Entity:     NewEntity(uuid.New().String(), name),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}
	return evt
}

func (e event) EventName() string     { return e.name }
func (e event) Payload() EventPayload { return e.payload }
func (e event) Metadata() Metadata    { return e.metadata }
func (e event) OccurredAt() time.Time { return e.occurredAt }
