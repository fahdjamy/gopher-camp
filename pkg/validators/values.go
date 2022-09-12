package validators

import "reflect"

func IsIntEmpty(val int) bool {
	return reflect.ValueOf(val).IsZero()
}
