package store

import (
	"encoding/json"
	"fmt"

	"github.com/amebalabs/nocreep/app/model"
	bolt "github.com/coreos/bbolt"
)

// BoltDB db holder
type BoltDB struct {
	db *bolt.DB
}

const (
	deviceBucketName = "device"
	eventBucketName  = "event"
)

// SetupBoltDB Inits bolt database
func SetupBoltDB() (*BoltDB, error) {
	result := BoltDB{}
	db, err := bolt.Open("nocreep.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	topBuckets := []string{deviceBucketName, eventBucketName}

	err = db.Update(func(tx *bolt.Tx) error {
		for _, bktName := range topBuckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(bktName)); e != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	result.db = db
	return &result, nil
}

//Close database
func (b *BoltDB) Close() error {
	return b.db.Close()
}

func (b *BoltDB) save(bkt *bolt.Bucket, key []byte, value interface{}) (err error) {
	if value == nil {
		return fmt.Errorf("can't save nil value for %s", key)
	}
	jdata, jerr := json.Marshal(value)
	if jerr != nil {
		return fmt.Errorf("can't marshal comment: %s", jerr)
	}
	if err = bkt.Put(key, jdata); err != nil {
		return fmt.Errorf("failed to save key %s", key)
	}
	return nil
}

func (b *BoltDB) load(bkt *bolt.Bucket, key []byte, res interface{}) error {
	value := bkt.Get(key)
	if value == nil {
		return fmt.Errorf("no value for %s", key)
	}

	if err := json.Unmarshal(value, &res); err != nil {
		return fmt.Errorf("failed to unmarshal: %s", err)
	}
	return nil
}

//AddUser TODO
func (b *BoltDB) AddUser(model.User) (err error) {
	return err
}

//AddDevice adds device to store
func (b *BoltDB) AddDevice(device model.Device) (err error) {

	err = b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(deviceBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", deviceBucketName)
		}

		if e := b.save(bucket, []byte(device.ID), device); e != nil {
			return fmt.Errorf("failed to put key %s to bucket %s", device.ID, deviceBucketName)
		}

		return nil
	})

	return err
}

//GetDevices lookup device by it's id
func (b *BoltDB) GetDevices(model.User) (devices []model.Device, err error) {

	err = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(deviceBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", deviceBucketName)
		}

		return bucket.ForEach(func(k, v []byte) error {
			device := model.Device{}
			if e := json.Unmarshal(v, &device); e != nil {
				return fmt.Errorf("failed to unmarshal: %q", e)
			}
			devices = append(devices, device)
			return nil
		})
	})

	return devices, err
}

// getDeviceBucket return bucket with all events for device
func (b *BoltDB) getDeviceBucket(tx *bolt.Tx, deviceID model.DeviceID) (*bolt.Bucket, error) {
	eventBucket := tx.Bucket([]byte(eventBucketName))
	if eventBucket == nil {
		return nil, fmt.Errorf("no bucket %s", eventBucketName)
	}
	res := eventBucket.Bucket([]byte(deviceID))
	if res == nil {
		return nil, fmt.Errorf("no bucket %s in store", deviceID)
	}
	return res, nil
}

// makeDeviceBucket create new bucket for deviceID as a key. This bucket holds all events for device
func (b *BoltDB) makeDeviceBucket(tx *bolt.Tx, deviceID model.DeviceID) (*bolt.Bucket, error) {
	eventBucket := tx.Bucket([]byte(eventBucketName))
	if eventBucket == nil {
		return nil, fmt.Errorf("no bucket %s", eventBucketName)
	}
	res, err := eventBucket.CreateBucketIfNotExists([]byte(deviceID))
	if err != nil {
		return nil, fmt.Errorf("no bucket %s in store", deviceID)
	}
	return res, nil
}

//RecordEvent records event
func (b *BoltDB) RecordEvent(id model.DeviceID, event model.Event) (err error) {
	err = b.db.Update(func(tx *bolt.Tx) error {
		bucket, e := b.makeDeviceBucket(tx, id)
		if e != nil {
			return e
		}
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", eventBucketName)
		}

		if bucket.Get([]byte(event.ID)) != nil {
			return fmt.Errorf("key %s already in store", event.ID)
		}

		if e := b.save(bucket, []byte(event.ID), event); e != nil {
			return fmt.Errorf("failed to put key %s to bucket %s", event.ID, eventBucketName)
		}

		return nil
	})
	return err
}

//RecordEvents records event
func (b *BoltDB) RecordEvents(model.DeviceID, []model.Event) (err error) {
	return err
}

//GetDeviceEvents retrives all events for provided Device
func (b *BoltDB) GetDeviceEvents(id model.DeviceID) (events []model.Event, err error) {
	err = b.db.View(func(tx *bolt.Tx) error {
		bucket, e := b.getDeviceBucket(tx, id)
		if e != nil {
			return e
		}

		return bucket.ForEach(func(k, v []byte) error {
			event := model.Event{}
			if e := json.Unmarshal(v, &event); e != nil {
				return fmt.Errorf("failed to unmarshal: %q", e)
			}
			events = append(events, event)
			return nil
		})
	})

	return events, err
}

//GetUserEvents TODO
func (b *BoltDB) GetUserEvents(model.User) (events []model.Event, err error) {
	return events, err
}

//StopCollecting TODO
func (b *BoltDB) StopCollecting(model.DeviceID) (err error) {
	return err
}
