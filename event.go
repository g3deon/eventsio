package eventsio

import (
	"fmt"
	"sync"
	"time"

	"github.com/goccy/go-json"
)

type Event interface {
	GetID() string
	GetSendAt() time.Time
	GetHeader(key string) any
	SetHeader(key string, value any)
	GetHeaders() map[string]any
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type BaseEvent struct {
	ID      string         `json:"id"`
	SendAt  time.Time      `json:"sendAt"`
	Headers map[string]any `json:"headers"`
	mu      sync.RWMutex
}

func NewEvent(id string) *BaseEvent {
	return &BaseEvent{
		ID:     id,
		SendAt: time.Now().UTC(),
	}
}

func (e *BaseEvent) SetID(id string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ID = id
}

func (e *BaseEvent) GetID() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.ID
}

func (e *BaseEvent) SetSendAt(sendAt time.Time) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.SendAt = sendAt
}

func (e *BaseEvent) GetSendAt() time.Time {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.SendAt
}

func (e *BaseEvent) SetHeader(key string, value any) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Headers[key] = value
}

func (e *BaseEvent) GetHeader(key string) any {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.Headers[key]
}

func (e *BaseEvent) GetHeaders() map[string]any {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.Headers
}

func (e *BaseEvent) String() string {
	return fmt.Sprintf("BaseEvent{id: %s, sendAt: %s}", e.ID, e.SendAt)
}

func (e *BaseEvent) Marshal() ([]byte, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return json.Marshal(e)
}

func (e *BaseEvent) Unmarshal(data []byte) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	return json.Unmarshal(data, e)
}
