package store

//Generic interface for storage engine
import (
	"github.com/amebalabs/nocreep/app/model"
)

// SourceDataProvider methods to write collected analytics data
type SourceDataProvider interface {
	AddDevice(model.Device) error
}

// DataExplorer methods to access analytics data
type DataExplorer interface {
	GetDevices() ([]model.Device, error)
}
