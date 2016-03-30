package refl

import (
	"reflect"

	"github.com/tolexo/aero/panik"
)

func IsAddress(addr interface{}) bool {
	rt := reflect.TypeOf(addr)
	return rt.Kind() == reflect.Ptr
}

func ComposedOf(item interface{}, parent interface{}) bool {

	it := reflect.TypeOf(item)
	if it.Kind() == reflect.Ptr {
		it = it.Elem()
	}

	pt := reflect.TypeOf(parent)
	if pt.Kind() == reflect.Ptr {
		pt = pt.Elem()
	}
	panik.If(pt.Kind() != reflect.Struct, "parent must be struct type")

	// find field with parent's exact name
	f, ok := it.FieldByName(pt.Name())
	if !ok {
		return false
	}

	if !f.Anonymous {
		return false
	}

	if !f.Type.ConvertibleTo(pt) {
		return false
	}

	return true
}
