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
)

// SetupBoltDB Inits bolt database
func SetupBoltDB() (*BoltDB, error) {
	result := BoltDB{}
	db, err := bolt.Open("nocreep.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	topBuckets := []string{deviceBucketName}

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

//GetDevice lookup device by it's id
func (b *BoltDB) GetDevice() (devices []model.Device, err error) {

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
