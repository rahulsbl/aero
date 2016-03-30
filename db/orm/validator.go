package orm

import (
	"fmt"
	"reflect"

	"github.com/thejackrabbit/aero/str"
)

// insert and update values: must, may (default), no

func Insertable(modl interface{}, data map[string]string) (bool, []error) {

	success := true
	var errs []error

	mt := reflect.TypeOf(modl)
	if mt.Kind() == reflect.Ptr {
		mt = mt.Elem()
	}
	input := clone(data)

	for i := 0; i < mt.NumField(); i++ {
		fld := mt.Field(i)
		name := fld.Name
		sql := str.SnakeCase(name)
		_, ok := input[sql]

		if ok == false && fld.Tag.Get("insert") == "must" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Compulsory field missing: %s", sql))
		}

		if ok == true && fld.Tag.Get("insert") == "no" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Unneeded field present: %s", sql))
		}
	}

	return success, errs
}

func Updatable(modl interface{}, data map[string]string) (bool, []error) {
	success := true
	var errs []error

	mt := reflect.TypeOf(modl)
	if mt.Kind() == reflect.Ptr {
		mt = mt.Elem()
	}
	input := clone(data)

	for i := 0; i < mt.NumField(); i++ {
		fld := mt.Field(i)
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
			errs = append(errs, fmt.Errorf("Unneeded field present: %s", sql))
		}
	}

	return success, errs
}

func modelType(modl interface{}) reflect.Type {
	mt := reflect.TypeOf(modl)
	if mt.Kind() == reflect.Ptr {
		return mt.Elem()
	}
	return mt
}

func clone(data map[string]string) map[string]interface{} {
	out := make(map[string]interface{})
	for key, val := range data {
		out[key] = val
	}
	return out
}
