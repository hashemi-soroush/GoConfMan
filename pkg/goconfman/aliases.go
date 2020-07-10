package goconfman

import (
	"reflect"
)

type ConfigWithAliases interface {
	BindAliases()
}

func LoadFromAliases(config interface{}) {
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
			LoadFromAliases(fieldVal.Interface())
		}
	}

	c, ok := config.(ConfigWithAliases)
	if ok {
		c.BindAliases()
	}
}
