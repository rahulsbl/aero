package orm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/thejackrabbit/aero/str"
)

func Insertable(modl interface{}, data map[string]string) (bool, []error) {

	success := true
	var errs []error

	modlType := modelType(modl)
	fields := fields(modlType)
	input := clone(data)

	for i := 0; i < len(fields); i++ {
		fld := fields[i]
		_, ok := input[fld.sql]

		if ok == false && fld.insert == "must" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Compulsory field missing: %s", fld.sql))
		}

		if ok == true && fld.insert == "no" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Unneeded field present: %s", fld.sql))
		}
	}

	return success, errs
}

func Updatable(modl interface{}, data map[string]string) (bool, []error) {
	success := true
	var errs []error

	modlType := modelType(modl)
	fields := fields(modlType)
	input := clone(data)

	for i := 0; i < len(fields); i++ {
		fld := fields[i]
		_, ok := input[fld.sql]

		if ok == false && fld.update == "must" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Compulsory field missing: %s", fld.sql))
		}

		if ok == true && fld.update == "no" {
			success = false
			if errs == nil {
				errs = make([]error, 0)
			}
			errs = append(errs, fmt.Errorf("Unneeded field present: %s", fld.sql))
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

type field struct {
	name   string
	sql    string
	insert string
	update string
}

// insert and update values: must, may (default), no

func fields(mType reflect.Type) []field {
	flds := make([]field, mType.NumField())
	for f := 0; f < len(flds); f++ {
		fld := mType.Field(f)
		flds[f] = field{
			name:   fld.Name,
			sql:    str.SnakeCase(fld.Name),
			insert: strings.ToLower(fld.Tag.Get("insert")),
			update: strings.ToLower(fld.Tag.Get("update")),
		}
	}
	return flds
}

func clone(data map[string]string) map[string]interface{} {
	out := make(map[string]interface{})
	for key, val := range data {
		out[key] = val
	}
	return out
}
