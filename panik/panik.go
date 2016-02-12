package panik

import (
	"fmt"
)

func On(err error) {
	if err != nil {
		panic(err)
	}
}

func On2(err error, f func()) {
	if err != nil {
		if f != nil {
			f()
		}
		panic(err)
	}
}

func If(condition bool, message string, params ...interface{}) {
	if condition {
		s := fmt.Sprintf(message, params...)
		panic(s)
	}
}

func If2(condition bool, f func(), message string, params ...interface{}) {
	if condition {
		if f != nil {
			f()
		}
		s := fmt.Sprintf(message, params...)
		panic(s)
	}
}

func Do(message string, params ...interface{}) {
	If(true, message, params...)
}
