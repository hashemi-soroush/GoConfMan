package goconfman

import (
	"reflect"
)

// TODO: this interface should be declared on Config not *Config. why???
type ConfigWithDefaults interface {
	BindDefaults()
}

// TODO: this method should be able to access unexported fields
func LoadFromDefaults(config ConfigWithDefaults) {
	config.BindDefaults()

	configVal := reflect.ValueOf(config).Elem()
	for i := 0 ; i < configVal.NumField() ; i++ {
		fieldVal := configVal.Field(i)

		if fieldVal.Kind() == reflect.Struct {
			fieldVal = fieldVal.Addr()
		}

		if fieldVal.Kind() == reflect.Ptr {
			innerConfig, ok := fieldVal.Interface().(ConfigWithDefaults)
			if ok {
				LoadFromDefaults(innerConfig)
			}
		}
	}
}
