package goconfman

import (
	"reflect"
)

type ConfigWithAliases interface {
	BindAliases()
}

func LoadFromAliases(config ConfigWithAliases) {
	configVal := reflect.ValueOf(config).Elem()
	for i := 0 ; i < configVal.NumField() ; i++ {
		fieldVal := configVal.Field(i)

		if fieldVal.Kind() == reflect.Struct {
			fieldVal = fieldVal.Addr()
		}

		if fieldVal.Kind() == reflect.Ptr {
			innerConfig, ok := fieldVal.Interface().(ConfigWithAliases)
			if ok {
				LoadFromAliases(innerConfig)
			}
		}
	}

	config.BindAliases()
}
