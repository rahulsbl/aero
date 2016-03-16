package db

import (
	"database/sql/driver"
	"encoding/json"
)

// Add support for Json fields
// http://www.booneputney.com/2015-06-18-gorm-golang-jsonb-value-copy/

type JMap map[string]interface{}

func NewJMap() *JMap {
	return new(JMap)
}

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

func (j *JMap) Set(key string, val interface{}) {
	(*j)[key] = val
}

type JArr []interface{}

func NewJArr() *JArr {
	return new(JArr)
}

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
