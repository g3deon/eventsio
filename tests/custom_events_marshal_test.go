package eventsio_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CustomEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Struct    struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	} `json:"struct"`
	Array []string               `json:"array"`
	Map   map[string]interface{} `json:"map"`
	Bool  bool                   `json:"bool"`
	Int   int                    `json:"int"`
	Float float64                `json:"float"`
	Null  interface{}            `json:"null"`
}

func TestCustomEventMarshal(t *testing.T) {
	tests := []struct {
		name        string
		event       CustomEvent
		expectedErr bool
		validate    func(t *testing.T, jsonData []byte)
	}{
		{
			name: "Marshaling complete custom event",
			event: func() CustomEvent {
				e := CustomEvent{
					Timestamp: time.Date(2023, 10, 15, 12, 30, 0, 0, time.UTC),
					Bool:      true,
					Int:       42,
					Float:     3.14159,
					Null:      nil,
					Array:     []string{"one", "two", "three"},
					Map: map[string]interface{}{
						"key1": "value1",
						"key2": 2,
					},
				}
				e.Struct.Field1 = "Hello World"
				e.Struct.Field2 = 100
				return e
			}(),
			expectedErr: false,
			validate: func(t *testing.T, jsonData []byte) {
				var data map[string]any
				err := json.Unmarshal(jsonData, &data)
				assert.NoError(t, err, "Generated JSON should be valid")

				fields := []string{"timestamp", "struct", "array", "map", "bool", "int", "float", "null"}
				for _, field := range fields {
					assert.Contains(t, data, field, "JSON should contain field '%s'", field)
				}

				assert.True(t, data["bool"].(bool), "Field 'bool' should be true")
				assert.Equal(t, float64(42), data["int"].(float64), "Field 'int' should be 42")

				structData := data["struct"].(map[string]interface{})
				assert.Equal(t, "Hello World", structData["field1"], "Field 'struct.field1' should be 'Hello World'")
			},
		},
		{
			name: "Marshaling with empty arrays and maps",
			event: func() CustomEvent {
				e := CustomEvent{
					Timestamp: time.Date(2023, 10, 15, 12, 30, 0, 0, time.UTC),
					Array:     []string{},
					Map:       map[string]interface{}{},
				}
				return e
			}(),
			expectedErr: false,
			validate: func(t *testing.T, jsonData []byte) {
				var data map[string]interface{}
				err := json.Unmarshal(jsonData, &data)
				assert.NoError(t, err, "Generated JSON should be valid")

				array := data["array"].([]interface{})
				assert.Empty(t, array, "Array should be empty")

				mapData := data["map"].(map[string]interface{})
				assert.Empty(t, mapData, "Map should be empty")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tt.event)

			if tt.expectedErr {
				assert.Error(t, err, "Expected error but got none")
				return
			}

			assert.NoError(t, err, "Unexpected error during marshaling")
			if tt.validate != nil {
				tt.validate(t, jsonData)
			}
		})
	}
}

func TestCustomEventUnmarshal(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    string
		expectedErr bool
		validate    func(t *testing.T, event CustomEvent)
	}{
		{
			name: "Unmarshal complete JSON to CustomEvent",
			jsonData: `{
				"timestamp": "2023-10-15T12:30:00Z",
				"struct": {
					"field1": "Hello World",
					"field2": 100
				},
				"array": ["one", "two", "three"],
				"map": {
					"key1": "value1",
					"key2": 2
				},
				"bool": true,
				"int": 42,
				"float": 3.14159,
				"null": null
			}`,
			expectedErr: false,
			validate: func(t *testing.T, event CustomEvent) {
				expectedTime := time.Date(2023, 10, 15, 12, 30, 0, 0, time.UTC)
				assert.True(t, event.Timestamp.Equal(expectedTime), "Timestamp mismatch")

				assert.Equal(t, "Hello World", event.Struct.Field1, "Field1 mismatch")
				assert.Equal(t, 100, event.Struct.Field2, "Field2 mismatch")

				expectedArray := []string{"one", "two", "three"}
				assert.Equal(t, expectedArray, event.Array, "Array mismatch")

				assert.Equal(t, "value1", event.Map["key1"], "Map[key1] mismatch")

				assert.True(t, event.Bool, "Bool should be true")
				assert.Equal(t, 42, event.Int, "Int mismatch")
				assert.Equal(t, 3.14159, event.Float, "Float mismatch")
				assert.Nil(t, event.Null, "Null should be nil")
			},
		},
		{
			name: "Unmarshal JSON with missing fields",
			jsonData: `{
				"timestamp": "2023-10-15T12:30:00Z",
				"bool": true,
				"int": 42
			}`,
			expectedErr: false,
			validate: func(t *testing.T, event CustomEvent) {
				expectedTime := time.Date(2023, 10, 15, 12, 30, 0, 0, time.UTC)
				assert.True(t, event.Timestamp.Equal(expectedTime), "Timestamp mismatch")

				assert.True(t, event.Bool, "Bool should be true")
				assert.Equal(t, 42, event.Int, "Int mismatch")

				assert.Empty(t, event.Struct.Field1, "Field1 should be empty")
				assert.Equal(t, 0, event.Struct.Field2, "Field2 should be 0")
				assert.Empty(t, event.Array, "Array should be empty")
			},
		},
		{
			name: "Unmarshal invalid JSON",
			jsonData: `{
				"timestamp": "2023-10-15T12:30:00Z",
				"invalid json format
			}`,
			expectedErr: true,
			validate:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event CustomEvent
			err := json.Unmarshal([]byte(tt.jsonData), &event)

			if tt.expectedErr {
				assert.Error(t, err, "Expected error but got none")
				return
			}

			assert.NoError(t, err, "Unexpected error during unmarshaling")
			if tt.validate != nil {
				tt.validate(t, event)
			}
		})
	}
}
