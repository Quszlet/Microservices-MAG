package kafka

import (
	"encoding/json"
)

// Event describes a message exchanged via Kafka between services.
type Event struct {
	EventType string          `json:"event_type"`
	Timestamp int64           `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

// DecodeData unmarshals event payload to provided destination.
func (e *Event) DecodeData(dst interface{}) error {
	return json.Unmarshal(e.Data, dst)
}
