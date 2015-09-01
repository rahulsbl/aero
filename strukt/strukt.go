package strukt

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

func FromJson(addr interface{}, jsn string) error {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsn), addr)
	if err != nil {
		return err
	}
	return FromMap(addr, data)
}

func FromMap(addr interface{}, data map[string]interface{}) error {

	// TODO: use same decoder for all objects of same type?
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		Result:           addr,
		WeaklyTypedInput: false})

	if err != nil {
		return err
	}

	err = decoder.Decode(data)
	return err
}

func ToJson(s interface{}) (string, error) {
	b, e := json.Marshal(s)
	if e != nil {
		return "", e
	} else {
		return string(b), nil
	}
}

func ToBytes(s interface{}) ([]byte, error) {
	return json.Marshal(s)
}

func ToBytesPretty(s interface{}, pretty bool) ([]byte, error) {
	if !pretty {
		return json.Marshal(s)
	} else {
		return json.MarshalIndent(s, "", "\t")
	}
}
