package eventsio_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/g3deon/eventsio"
	"github.com/stretchr/testify/assert"
)

func TestEventMarshal(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		wantJSON string
	}{
		{
			name:     "marshal event with id",
			id:       "test-event-id",
			wantJSON: `{"id":"test-event-id"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			event := eventsio.NewEvent(test.id)

			data, err := event.Marshal()
			assert.NoError(t, err)

			var eventMap map[string]interface{}
			err = json.Unmarshal(data, &eventMap)
			assert.NoError(t, err)
			assert.Equal(t, test.id, eventMap["id"])
		})
	}
}

func TestEventUnmarshal(t *testing.T) {
	tests := []struct {
		name       string
		jsonData   string
		wantID     string
		wantSendAt time.Time
		wantError  bool
	}{
		{
			name:       "unmarshal event with id",
			jsonData:   `{"id":"test-event-id","sendAt":"2023-01-01T12:00:00Z"}`,
			wantID:     "test-event-id",
			wantSendAt: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			wantError:  false,
		},
		{
			name:      "unmarshal invalid json",
			jsonData:  `{invalid json}`,
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			event := &eventsio.BaseEvent{}
			err := event.Unmarshal([]byte(test.jsonData))

			event.SetSendAt(test.wantSendAt)

			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.wantID, event.GetID())
				if !test.wantError {
					assert.Equal(t, test.wantSendAt, event.GetSendAt())
				}
			}
		})
	}
}
