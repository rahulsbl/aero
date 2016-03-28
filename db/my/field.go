package my

import (
	"fmt"
	"reflect"
)

type field struct {
	Field   string  `gorm:"column:Field"`
	Type    string  `gorm:"column:Type"`
	Null    string  `gorm:"column:Null"`
	Key     string  `gorm:"column:Key"`
	Default *string `gorm:"column:Default"`
	Extra   string  `gorm:"column:Extra"`
}

func (f *field) info() string {
	key := fmt.Sprintf("%s %s", f.Field, f.Type)
	if f.Null == "NO" {
		key += " NOT NULL"
	}
	if f.Default != nil {
		key += " DEFAULT " + *(f.Default)
	}
	key += " " + f.Extra
	return key
}

func NestedStructFields(ifc interface{}) []reflect.StructField {
	fields := make([]reflect.StructField, 0)

	ift := reflect.TypeOf(ifc)
	ifv := reflect.ValueOf(ifc)

	for i := 0; i < ift.NumField(); i++ {

		fv := ifv.Field(i)
		ft := ift.Field(i)

		if fv.Kind() == reflect.Struct {
			fields = append(fields, NestedStructFields(fv.Interface())...)
		} else {
			fields = append(fields, ft)
		}
	}

	return fields
}
