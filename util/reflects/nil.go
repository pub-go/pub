package reflects

import "reflect"

func IsNil(input interface{}) bool {
	if input == nil {
		return true
	}
	rfValue := reflect.ValueOf(input)
	kind := rfValue.Kind()
	nullable := kind == reflect.Slice ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.Ptr ||
		kind == reflect.Map
	return nullable && rfValue.IsNil()
}
