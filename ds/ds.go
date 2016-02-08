// misc data-structure utility functions
package ds

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/thejackrabbit/aero/panik"
)

func Load(addr interface{}, b []byte) error {
	return json.Unmarshal(b, addr)
}

func Load2(addr interface{}, b []byte, enc string) error {
	switch enc {
	case "json":
		return json.Unmarshal(b, addr)
	case "ffjson":
		return ffjson.Unmarshal(b, addr)
	default:
		panik.Do("Unknown decoding %s", enc)
		return errors.New("Unknown decoding " + enc)
	}
}

func LoadStruct(addr interface{}, m map[string]interface{}) error {
	// TODO: use same decoder for all objects of same type?
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		Result:           addr,
		WeaklyTypedInput: false})

	if err != nil {
		return err
	}

	return decoder.Decode(m)
}

// Encoding to bytes
// -------------------

func ToBytes(o interface{}, jsonPretty bool) ([]byte, error) {
	if !jsonPretty {
		return json.Marshal(o)
	} else {
		return json.MarshalIndent(o, "", "\t")
	}
}

func ToBytes2(o interface{}, enc string) ([]byte, error) {
	switch enc {
	case "json":
		return json.Marshal(o)
	case "ffjson":
		return ffjson.Marshal(o)
	default:
		panik.Do("Unknown encoding %s", enc)
		return nil, errors.New("Unknown encoding " + enc)
	}
}
