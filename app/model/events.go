package model

// Event related info
type Event struct {
	ID       string   `json:"id"`
	DeviceID DeviceID `json:"os"`
}
