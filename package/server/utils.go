package server

import "reflect"

func is_nil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

func not_nil(a interface{}) bool {
	return !is_nil(a)
}
