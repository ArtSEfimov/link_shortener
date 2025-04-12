package event

type Event struct {
	Type string
	Data any
}

type Bus struct {
	bus chan Event
}

func NewBus() *Bus {
	return &Bus{
		bus: make(chan Event),
	}
}

func (bus *Bus) Publish(event Event) {
	bus.bus <- event
}

func (bus *Bus) Subscribe() <-chan Event {
	return bus.bus
}
