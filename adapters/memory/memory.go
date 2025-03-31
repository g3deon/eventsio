package memory

import (
	"sync"

	"github.com/g3deon/eventsio"
)

type EventBus struct {
	subscribers map[string][]func(eventsio.Event)
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(eventsio.Event)),
	}
}

func (b *EventBus) Publish(topic string, event eventsio.Event) error {
	b.mu.RLock()
	handlers := b.subscribers[topic]
	b.mu.RUnlock()

	for _, handler := range handlers {
		handler(event)
	}
	return nil
}

func (b *EventBus) Subscribe(topic string, handler func(eventsio.Event)) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[topic] = append(b.subscribers[topic], handler)
	return nil
}
