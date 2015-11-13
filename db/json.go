package db

import (
	"database/sql/driver"
	"encoding/json"
)

// Add support for jsonb fields
// http://www.booneputney.com/2015-06-18-gorm-golang-jsonb-value-copy/

type JMap map[string]interface{}

func (j JMap) Value() (driver.Value, error) {
	str, err := json.Marshal(j)
	return string(str), err
}

func (j *JMap) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

// TODO:
type JSlice []interface{}
