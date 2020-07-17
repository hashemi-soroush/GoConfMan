package goconfman

import (
	"fmt"
	"os"
	"reflect"
)

type ConfigWithEnvVars interface {
	BindEnvVars(prefix string)
}

func LoadFromEnvVars(config interface{}, prefix string) {
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
			newPrefix := fieldVal.Type().Elem().Name()
			if prefix != "" {
				newPrefix = fmt.Sprintf("%s__%s", prefix, newPrefix)
			}
			LoadFromEnvVars(fieldVal.Interface(), newPrefix)
		}
	}

	c, ok := config.(ConfigWithEnvVars)
	if ok {
		c.BindEnvVars(prefix)
	}
}

func BindEnvVar(field interface{}, envVarName string, prefix string) {
	fieldValue := reflect.ValueOf(field)
	if fieldValue.Kind() != reflect.Ptr {
		panic("goconfman.BindEnvVar's field argument must a pointer, so goconfman can change its value")
	}

	if prefix != "" {
		envVarName = fmt.Sprintf("%s__%s", prefix, envVarName)
	}
	valueString, ok := os.LookupEnv(envVarName)
	if ok == false {
		return
	}

	err := genericSet(field, valueString)
	if err != nil {
		panic(fmt.Sprintf("Failed to load env var %s\n%s", envVarName, err))
	}
}
