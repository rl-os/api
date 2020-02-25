package utils

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type NullString struct {
	sql.NullString
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

func ValueOfStruct(dest interface{}) (driver.Value, error) {
	return json.Marshal(dest)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func ScanToStruct(dest interface{}, value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	return mapstructure.Decode(data, dest)
}
