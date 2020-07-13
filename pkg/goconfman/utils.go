package goconfman

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

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

func genericSet(field interface{}, valueString string) error {
	fieldValue := reflect.ValueOf(field).Elem()
	fieldType := fieldValue.Type()
	switch field.(type) {
	case *string:
		fieldValue.SetString(valueString)
	case *int, *int8, *int16, *int32, *int64:
		parsedValue, err := strconv.ParseInt(valueString, 0, fieldType.Bits())
		if err != nil {
			return errors.New(fmt.Sprintf("\"%s\" can't be parsed into a %s because: %s", valueString, fieldType.Kind().String(), err))
		}
		fieldValue.SetInt(parsedValue)
	case *float32, *float64:
		parsedValue, err := strconv.ParseFloat(valueString, fieldType.Bits())
		if err != nil {
			return errors.New(fmt.Sprintf("\"%s\" can't be parsed into a %s because: %s", valueString, fieldType.Kind().String(), err))
		}
		fieldValue.SetFloat(parsedValue)
	case bool:
		parsedValue, err := strconv.ParseBool(valueString)
		if err != nil {
			return errors.New(fmt.Sprintf("\"%s\" can't be parsed into a %s because: %s", valueString, fieldType.Kind().String(), err))
		}
		fieldValue.SetBool(parsedValue)
	default:
		return errors.New("TODO")
	}

	return nil
}