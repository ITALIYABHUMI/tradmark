package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// MarshalJSONHelper marshals a struct to JSON.
func MarshalJSONHelper(v interface{}) (driver.Value, error) {
	return json.Marshal(v)
}

// UnmarshalJSONHelper unmarshals JSON data into a struct.
func UnmarshalJSONHelper(value interface{}, dest interface{}) error {
	if value == nil {
		return errors.New("nil value")
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan")
	}


	return json.Unmarshal(bytes, dest)
}
