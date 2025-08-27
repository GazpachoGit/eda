package consumer

import (
	"context"
	"domain-event/internal/am"
	"domain-event/internal/ddd"
	"domain-event/internal/domain"
	"domain-event/internal/handlers"
	"domain-event/internal/jetstream"
	"domain-event/internal/registry"
	"domain-event/internal/waiter"
	"fmt"

	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
)

const (
	cfg_Nats_URL    = ""
	cfg_Nats_Stream = ""
)

type App struct {
	nc               *nats.Conn
	js               nats.JetStreamContext
	waiter           waiter.Waiter
	domainDispatcher *ddd.EventDispatcher
}

func NewApp() (*App, error) {
	nc, err := nats.Connect(cfg_Nats_URL)
	if err != nil {
		return nil, err
	}
	defer nc.Close()
	js, err := initJetStream(nc)
	if err != nil {
		return nil, err
	}

	waiter := waiter.New(waiter.CatchSignals())

	eventStream := am.NewEventStream(registry.NewRegistry(), jetstream.NewStream(cfg_Nats_Stream, js))
	domainDispatcher := ddd.NewEventDispatcher()
	consumerHandler := handlers.NewConsumedEventHandler()
	handlers.RegisterConsumerEventHandler(eventStream, consumerHandler.HandleEvent)

	a := &App{
		nc, js, waiter, domainDispatcher,
	}

	a.waiter.Add(
		a.waitForStream,
	)
	return a, nil
}

func initJetStream(nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     cfg_Nats_Stream,
		Subjects: []string{fmt.Sprintf("%s.>", cfg_Nats_Stream)},
	})

	return js, err
}

func (a *App) waitForStream(ctx context.Context) error {
	closed := make(chan struct{})
	a.nc.SetClosedHandler(func(*nats.Conn) {
		close(closed)
	})
	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("message stream started")
		defer fmt.Println("message stream stopped")
		<-closed
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		return a.nc.Drain()
	})
	return group.Wait()
}

func (a *App) Waiter() waiter.Waiter {
	return a.waiter
}

func (a *App) SendStoreCreatedEvent(ID int, Name string, Location string) {
	event := ddd.NewEvent(domain.StoreCreatedEvent, domain.StoreCreated{ID, Name, Location})
	a.domainDispatcher.Publish(context.Background(), event)
}
