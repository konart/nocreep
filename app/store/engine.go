package store

//Generic interface for storage engine
import (
	"github.com/amebalabs/nocreep/app/model"
)

// DataInterface ...
type DataInterface interface {
	DataProducer
	DataConsumer
}

// DataProducer methods to write collected analytics data
type DataProducer interface {
	//AddUser adds new device for analytics collection
	AddUser(model.User) error
	//AddDevice adds new device for analytics collection
	AddDevice(model.Device) error
	//RecordEvent records event for specified device
	RecordEvent(model.DeviceID, model.Event) error // probably unnecesary
	//RecordEvents records multiple events for specified device
	RecordEvents(model.DeviceID, []model.Event) error
}

// DataConsumer methods to access analytics data
type DataConsumer interface {
	//GetDevices lists devices for user
	GetDevices(model.User) ([]model.Device, error)
	//GetDeviceEvents lists events for device
	GetDeviceEvents(model.DeviceID) ([]model.Event, error)
	//GetUserEvents lists events for all user devices
	GetUserEvents(model.User) ([]model.Event, error)
	//StopCollecting stops events collection for specified device
	StopCollecting(model.DeviceID) error
}
