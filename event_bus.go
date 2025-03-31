package eventsio

type EventBus interface {
	Publish(topic string, event Event) error
	Subscribe(topic string, handler func(event Event)) error
}
