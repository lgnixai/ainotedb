
package events

type Event interface {
	EventName() string
	Payload() interface{}
}

type EventHandler interface {
	Handle(event Event) error
}

type EventBus interface {
	Publish(event Event) error
	Subscribe(handler EventHandler, events ...string)
}
