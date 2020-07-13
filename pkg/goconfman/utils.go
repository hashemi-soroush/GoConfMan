package goconfman

import "reflect"

func loadRecursive(config interface{}, loadFunc func(interface{})) {
	configVal := reflect.ValueOf(config)
	if configVal.Kind() == reflect.Ptr {
		configVal = configVal.Elem()
	}
	for i := 0; i < configVal.NumField(); i++ {
		fieldVal := configVal.Field(i)

		if fieldVal.Kind() == reflect.Struct {
			fieldVal = fieldVal.Addr()
		}

		if fieldVal.Kind() == reflect.Ptr {
			loadFunc(fieldVal.Interface())
		}
	}
}