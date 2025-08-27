package am

import (
	ddd "domain-event/internal/domain"
	"time"
)

type EventMessage interface {
	Message
	ddd.Event
}

type eventMessage struct {
	//ddd.Event:
	id         string
	name       string
	payload    ddd.EventPayload
	metadata   ddd.Metadata
	occurredAt time.Time
	//am.Message:
	//message from BUS(for incomming) and message to be sent(for outcomming)
	msg RawMessage
}

func (e eventMessage) ID() string                { return e.id }
func (e eventMessage) EventName() string         { return e.name }
func (e eventMessage) Payload() ddd.EventPayload { return e.payload }
func (e eventMessage) Metadata() ddd.Metadata    { return e.metadata }
func (e eventMessage) OccurredAt() time.Time     { return e.occurredAt }
func (e eventMessage) MessageName() string       { return e.msg.MessageName() }
func (e eventMessage) Ack() error                { return e.msg.Ack() }
func (e eventMessage) NAck() error               { return e.msg.NAck() }
func (e eventMessage) Extend() error             { return e.msg.Extend() }
func (e eventMessage) Kill() error               { return e.msg.Kill() }
