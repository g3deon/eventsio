package adapters_test

import (
	"testing"

	"github.com/g3deon/eventsio"
	"github.com/g3deon/eventsio/adapters/memory"
	"github.com/stretchr/testify/assert"
)

type testEvent struct {
	*eventsio.BaseEvent
	Data string
}

func (e *testEvent) Type() string {
	return "test_event"
}

func TestMemoryAdapter(t *testing.T) {
	tests := []struct {
		name     string
		topic    string
		event    eventsio.Event
		handlers []func(eventsio.Event)
		want     string
	}{
		{
			name:  "publish and receive simple event",
			topic: "test_topic",
			event: &testEvent{
				BaseEvent: eventsio.NewEvent("test-id"),
				Data:      "test_data",
			},
			handlers: []func(eventsio.Event){
				func(e eventsio.Event) {
					te := e.(*testEvent)
					assert.Equal(t, "test_data", te.Data)
				},
			},
		},
		{
			name:  "multiple handlers for the same topic",
			topic: "test_topic",
			event: &testEvent{
				BaseEvent: eventsio.NewEvent("test-id-2"),
				Data:      "test_data",
			},
			handlers: []func(eventsio.Event){
				func(e eventsio.Event) {
					te := e.(*testEvent)
					assert.Equal(t, "test_data", te.Data)
				},
				func(e eventsio.Event) {
					te := e.(*testEvent)
					assert.Equal(t, "test_data", te.Data)
				},
			},
		},
		{
			name:  "topic without handlers",
			topic: "empty_topic",
			event: &testEvent{
				BaseEvent: eventsio.NewEvent("test-id-3"),
				Data:      "test_data",
			},
			handlers: []func(eventsio.Event){},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := memory.NewEventBus()

			if len(tt.handlers) > 0 {
				for _, handler := range tt.handlers {
					err := bus.Subscribe(tt.topic, handler)
					assert.NoError(t, err)
				}
			}

			err := bus.Publish(tt.topic, tt.event)
			assert.NoError(t, err)
		})
	}
}

func TestMemoryAdapterConcurrency(t *testing.T) {
	bus := memory.NewEventBus()
	received := make(chan bool, 100)

	for range 10 {
		err := bus.Subscribe("concurrent_topic", func(e eventsio.Event) {
			received <- true
		})
		assert.NoError(t, err)
	}

	for range 10 {
		go func() {
			event := &testEvent{
				BaseEvent: eventsio.NewEvent("test-id"),
				Data:      "concurrent_test",
			}
			err := bus.Publish("concurrent_topic", event)
			assert.NoError(t, err)
		}()
	}

	for range 100 {
		<-received
	}
}
