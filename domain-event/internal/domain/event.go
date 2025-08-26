package ddd

import "time"

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
