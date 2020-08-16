package helper

import "reflect"

// IsPtr is pointer
func IsPtr(val interface{}) bool {
	v := reflect.ValueOf(val)
	return v.Kind() == reflect.Ptr
}

// IsArr is array
func IsArr(val interface{}) bool {
	v := reflect.ValueOf(val)
	return v.Kind() == reflect.Array
}
