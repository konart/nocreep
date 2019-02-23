package model

import (
	"time"
)

// Event related info
type Event struct {
	ID         string      `json:"id"`
	DeviceID   DeviceID    `json:"deviceid"`
	Name       string      `json:"name"`
	Type       string      `json:"type"` //TODO: Refactor to EventType
	Attributes interface{} `json:"attributes,omitempty"`
	RecorderAt time.Time   `json:"recordedAt,omitempty"` // when event happened
	CreatedAt  time.Time   `json:"createdAt,omitempty"`  // when it was delivered\stored on server side
}

// New event
func New(id string, deviceID DeviceID, name string) Event {
	return Event{ID: id, DeviceID: deviceID, Name: name, CreatedAt: time.Now()}
}
