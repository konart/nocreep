package store

//Generic interface for storage engine
import (
	"github.com/amebalabs/nocreep/app/model"
)

// SourceDataProvider methods to write collected analytics data
type SourceDataProvider interface {
	AddDevice(model.Device) error
	RecordEvent(model.DeviceID, model.Event) error
}

// DataExplorer methods to access analytics data
type DataExplorer interface {
	GetDevices() ([]model.Device, error)
	GetEvents(model.Device) ([]model.Event, error)
}
