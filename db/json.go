package db

import (
	"database/sql/driver"
	"encoding/json"
)

// Add support for Json fields
// http://www.booneputney.com/2015-06-18-gorm-golang-jsonb-value-copy/

type JDoc map[string]interface{}

func (j JDoc) Value() (driver.Value, error) {
	str, err := json.Marshal(j)
	return string(str), err
}

func (j *JDoc) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

type JArr []interface{}

func (j JArr) Value() (driver.Value, error) {
	str, err := json.Marshal(j)
	return string(str), err
}

func (j *JArr) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
