package validators

import "reflect"

func IsUintEmpty(val uint) bool {
	return reflect.ValueOf(val).IsZero()
}

func IsIntEmpty(val int) bool {
	return reflect.ValueOf(val).IsZero()
}
