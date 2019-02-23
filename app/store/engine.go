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
	AddDevice(model.Device) error
	RecordEvent(model.DeviceID, model.Event) error
}

// DataConsumer methods to access analytics data
type DataConsumer interface {
	GetDevices() ([]model.Device, error)
	GetEvents(model.DeviceID) ([]model.Event, error)
}
