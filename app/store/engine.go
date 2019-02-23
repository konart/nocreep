package store

//Generic interface for storage engine
import (
	"github.com/amebalabs/nocreep/app/model"
)

// SourceDataProvider methods to write collected analytics data
type SourceDataProvider interface {
}

// DataExplorer methods to access analytics data
type DataExplorer interface {
	GetDevice() ([]model.Device, error)
}
