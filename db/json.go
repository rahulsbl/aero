package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Add support for Json fields
// http://www.booneputney.com/2015-06-18-gorm-golang-jsonb-value-copy/

type JsonM map[string]interface{}

func NewJsonM() *JsonM {
	j := make(JsonM)
	return &j
}

func NewJsonM2(data map[string]interface{}) *JsonM {
	j := make(JsonM)
	for key := range data {
		j[key] = data[key]
	}
	return &j
}

func (j *JsonM) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	str, err := json.Marshal(j)
	return string(str), err
}

func (j *JsonM) Scan(value interface{}) error {
	if value == nil {
		j = nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}
	if err := json.Unmarshal(bytes, j); err != nil {
		return err
	}
	return nil
}

func (j *JsonM) Set(key string, val interface{}) *JsonM {
	(*j)[key] = val
	return j
}

type JsonA []interface{}

func NewJsonA(items ...interface{}) *JsonA {
	//return new(JsonA)
	len := len(items)
	arr := make(JsonA, len)
	for i := 0; i < len; i++ {
		arr[i] = items[i]
	}
	return &arr
}

func (j *JsonA) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	str, err := json.Marshal(j)
	return string(str), err
}

func (j *JsonA) Scan(value interface{}) error {
	if value == nil {
		j = nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}
	if err := json.Unmarshal(bytes, &j); err != nil {
		return err
	}
	return nil
}
