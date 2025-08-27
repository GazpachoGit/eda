package domain

const (
	StoreAggregateChannel = "mallbots.stores.events.Store"

	StoreCreatedEvent              = "stores.StoreCreated"
	StoreParticipatingToggledEvent = "stores.StoreParticipatingToggled"

	StoreCreatedEventAsync              = "storesapi.StoreCreated"
	StoreParticipatingToggledEventAsync = "storesapi.StoreParticipatingToggled"
)

type StoreCreated struct {
	ID       int
	Name     string
	Location string
}

type StoreParticipationToggled struct {
	ID            int
	Participating bool
}
