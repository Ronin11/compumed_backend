package device

import (
	"encoding/json"
	// "time"
	"errors"

	"database/sql/driver"
)

type Device struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

func (device *Device) Value() (driver.Value, error) {
	return json.Marshal(device)
}

func (device Device) Serialize() []byte {
	str, _ := json.Marshal(device)
	return str
}

func (device *Device) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &device)
}

// type DeviceAttributes struct {

// }

// func (attributes *DeviceAttributes) Value() (driver.Value, error) {
// 	return json.Marshal(attributes)
// }

// func (attributes DeviceAttributes) Serialize() []byte {
// 	str, _ := json.Marshal(attributes)
// 	return str
// }

// func (attributes *DeviceAttributes) Scan(value interface{}) error {
// 	b, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}

// 	return json.Unmarshal(b, &attributes)
// }
