package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/rightjoin/aero/ds"
)

// Add support for Json fields
// http://www.booneputney.com/2015-06-18-gorm-golang-jsonb-value-copy/

type JsonM map[string]interface{}

func NewJsonM2() *JsonM {
	j := make(JsonM)
	return &j
}

func NewJsonM(data map[string]interface{}) *JsonM {
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

type Json struct {
	Data interface{} `json:"data"`
}

func NewJson(data interface{}) *Json {
	return &Json{Data: data}
}

func NewJson2(str string) *Json {
	var d interface{}
	err := ds.Load(&d, []byte(str))
	if err != nil {
		panic(err)
	}
	return NewJson(d)
}

func (j *Json) Value() (driver.Value, error) {
	if j == nil || j.Data == nil {
		return nil, nil
	}
	str, err := json.Marshal(j.Data)
	return string(str), err
}

func (j *Json) Scan(value interface{}) error {

	if value == nil {
		j.Data = nil // j.Data or j (?)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}
	var load interface{}
	if err := json.Unmarshal(bytes, &load); err != nil {
		return err
	}
	j.Data = load
	return nil
}
