package model

// DeviceID type
type DeviceID string

// Device related info
type Device struct {
	ID          DeviceID `json:"id"`
	OS          string   `json:"os"`
	Jailbroken  bool     `json:"isJailbroken"`
	AppVersion  string   `json:"appVersion"`    //latest installed app version
	PushEnabled bool     `json:"isPushEnabled"` //Whether this user has opted in to receive Push Notifications
}
