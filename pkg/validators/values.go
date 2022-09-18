package validators

import "reflect"

func IsIntEmpty(val uint) bool {
	return reflect.ValueOf(val).IsZero()
}
