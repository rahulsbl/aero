package orm

import (
	"fmt"
	"reflect"

	"github.com/rightjoin/aero/ds"
	"github.com/rightjoin/aero/refl"
	"github.com/rightjoin/aero/str"
)

func Insertable(modl interface{}, data map[string]string) (bool, []error) {

	success := true
	var errs []error

	input := clone(data)

	obj := modl
	if reflect.TypeOf(modl).Kind() == reflect.Ptr {
		obj = reflect.ValueOf(modl).Elem()
	}

	for _, fld := range refl.NestedFields(obj) {
		name := fld.Name
		sql := str.SnakeCase(name)
		_, ok := input[sql]

		// must-insert validation
		if ok == false && fld.Tag.Get("insert") == "must" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Compulsory field missing: %s", sql))
		}

		// no-insert validation
		if ok == true && fld.Tag.Get("insert") == "no" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Extra field present: %s", sql))
		}

		// json_array and json_map validations
		sgnt := refl.TypeSignature(fld.Type)
		if ok {
			if sgnt == "sl:." || sgnt == "*sl:." {
				var test []interface{}
				if ds.Load(&test, []byte(data[sql])) != nil {
					success = false
					errs = append(errs, fmt.Errorf("Field must be json array: %s", sql))
				}
			} else if sgnt == "map" || sgnt == "*map" {
				var test map[string]interface{}
				if ds.Load(&test, []byte(data[sql])) != nil {
					success = false
					errs = append(errs, fmt.Errorf("Field must be json document: %s", sql))
				}
			} else if sgnt == "*sl:.uint8" || sgnt == "sl:.uint8" {
				var test interface{}
				if ds.Load(&test, []byte(data[sql])) != nil {
					success = false
					errs = append(errs, fmt.Errorf("Field must be valid json: %s", sql))
				}
			}
		}
	}

	return success, errs
}

func Updatable(modl interface{}, data map[string]string) (bool, []error) {
	success := true
	var errs []error

	input := clone(data)

	obj := modl
	if reflect.TypeOf(modl).Kind() == reflect.Ptr {
		obj = reflect.ValueOf(modl).Elem()
	}

	for _, fld := range refl.NestedFields(obj) {
		name := fld.Name
		sql := str.SnakeCase(name)
		_, ok := input[sql]

		if ok == false && fld.Tag.Get("update") == "must" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Compulsory field missing: %s", sql))
		}

		if ok == true && fld.Tag.Get("update") == "no" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Extra field present: %s", sql))
		}

		// json_array and json_map validations
		sgnt := refl.TypeSignature(fld.Type)
		if ok {
			if sgnt == "sl:." || sgnt == "*sl:." {
				var test []interface{}
				if ds.Load(&test, []byte(data[sql])) != nil {
					success = false
					errs = append(errs, fmt.Errorf("Field must be json array: %s", sql))
				}
			} else if sgnt == "map" || sgnt == "*map" {
				var test map[string]interface{}
				if ds.Load(&test, []byte(data[sql])) != nil {
					success = false
					errs = append(errs, fmt.Errorf("Field must be json document: %s", sql))
				}
			} else if sgnt == "*sl:.uint8" || sgnt == "sl:.uint8" {
				var test interface{}
				if ds.Load(&test, []byte(data[sql])) != nil {
					success = false
					errs = append(errs, fmt.Errorf("Field must be valid json: %s", sql))
				}
			}
		}
	}

	return success, errs
}

func clone(data map[string]string) map[string]interface{} {
	out := make(map[string]interface{})
	for key, val := range data {
		out[key] = val
	}
	return out
}
