package model

// DeviceID type
type DeviceID string

// Device related info
type Device struct {
	ID DeviceID `json:"id"`
	OS string   `json:"os"`
}
