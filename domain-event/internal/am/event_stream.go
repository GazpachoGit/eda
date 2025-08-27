package am

import (
	"context"
	ddd "domain-event/internal/ddd"
	"domain-event/internal/registry"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventPublisher = MessagePublisher[ddd.Event]
type EventSubscriber = MessageSubscriber[EventMessage]
type EventStream = MessageStream[ddd.Event, EventMessage]

type eventStream struct {
	reg    registry.Registry
	stream MessageStream[RawMessage, RawMessage]
}

func NewEventStream(reg registry.Registry, stream MessageStream[RawMessage, RawMessage]) EventStream {
	return &eventStream{
		reg:    reg,
		stream: stream,
	}
}

func (s eventStream) Publish(ctx context.Context, topicName string, event ddd.Event) error {
	//convert ddd.Event to EventMessage

	//covert metadata map to struct
	metadata, err := structpb.NewStruct(event.Metadata())
	if err != nil {
		return err
	}

	//Serialize domain event as the payload field
	payload, err := s.reg.Serialize(event.EventName(), event.Payload())
	if err != nil {
		return err
	}

	//prepare 'data' field using protobuf(data = payload + metadata)
	data, err := proto.Marshal(&EventMessageData{
		Payload:    payload,
		OccurredAt: timestamppb.New(event.OccurredAt()),
		Metadata:   metadata,
	})
	if err != nil {
		return err
	}

	localRawMessage := rawMessage{
		id:   event.ID(),
		name: event.EventName(),
		data: data,
	}

	//sent the EventMessage
	return s.stream.Publish(ctx, topicName, localRawMessage)
}

func (s eventStream) Subscribe(topicName string, handler MessageHandler[EventMessage], options ...SubscriberOption) error {
	cfg := NewSubscriberConfig(options)
	var filters map[string]struct{}
	if len(cfg.MessageFilters()) > 0 {
		filters = make(map[string]struct{})
		for _, key := range cfg.MessageFilters() {
			filters[key] = struct{}{}
		}
	}

	fn := MessageHandlerFunc[RawMessage](func(ctx context.Context, msg RawMessage) error {
		var eventData EventMessageData

		if filters != nil {
			if _, exists := filters[msg.MessageName()]; !exists {
				return nil
			}
		}

		err := proto.Unmarshal(msg.Data(), &eventData)
		if err != nil {
			return err
		}

		eventName := msg.MessageName()

		payload, err := s.reg.Deserialize(eventName, eventData.GetPayload())
		if err != nil {
			return err
		}

		eventMsg := eventMessage{
			id:         msg.ID(),
			name:       eventName,
			payload:    payload,
			metadata:   eventData.GetMetadata().AsMap(),
			occurredAt: eventData.GetOccurredAt().AsTime(),
			msg:        msg,
		}

		return handler.HandleMessage(ctx, eventMsg)
	})
	return s.stream.Subscribe(topicName, fn, options...)
}
