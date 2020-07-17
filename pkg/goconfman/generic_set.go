package goconfman

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func genericSet(field interface{}, stringValue string) error {
	fieldValue := reflect.ValueOf(field).Elem()

	var err error
	switch fieldValue.Type().Kind() {
	case reflect.String:
		err = setString(fieldValue, stringValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		err = setInts(fieldValue, stringValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		err = setUints(fieldValue, stringValue)
	case reflect.Float32, reflect.Float64:
		err = setFloats(fieldValue, stringValue)
	case reflect.Bool:
		err = setBool(fieldValue, stringValue)
	case reflect.Array, reflect.Slice:
		err = setSlices(fieldValue, stringValue)
	case reflect.Map:
		err = setMaps(fieldValue, stringValue)
	default:
		return errors.New(fmt.Sprintf("Type %s isn't supported by goconfman", fieldValue.Type().Kind().String()))
	}

	if err != nil {
		return errors.New(fmt.Sprintf("\"%s\" can't be parsed into %s because: %s", stringValue, fieldValue.Type().Kind().String(), err.Error()))
	}
	return nil
}

func setString(fieldValue reflect.Value, stringValue string) error {
	fieldValue.SetString(stringValue)
	return nil
}

func setInts(fieldValue reflect.Value, stringValue string) error {
	parsedValue, err := strconv.ParseInt(stringValue, 0, fieldValue.Type().Bits())
	if err != nil {
		return err
	}
	fieldValue.SetInt(parsedValue)
	return nil
}

func setUints(fieldValue reflect.Value, stringValue string) error {
	parsedValue, err := strconv.ParseUint(stringValue, 0, fieldValue.Type().Bits())
	if err != nil {
		return err
	}
	fieldValue.SetUint(parsedValue)
	return nil
}

func setFloats(fieldValue reflect.Value, stringValue string) error {
	parsedValue, err := strconv.ParseFloat(stringValue, fieldValue.Type().Bits())
	if err != nil {
		return err
	}
	fieldValue.SetFloat(parsedValue)
	return nil
}

func setBool(fieldValue reflect.Value, stringValue string) error {
	parsedValue, err := strconv.ParseBool(stringValue)
	if err != nil {
		return err
	}
	fieldValue.SetBool(parsedValue)
	return nil
}

func setSlices(fieldValue reflect.Value, stringValue string) error {
	slice := make([]interface{}, 0)
	err := json.Unmarshal([]byte(stringValue), &slice)
	if err != nil {
		return err
	}

	for _, element := range slice {
		elementBytes, err := json.Marshal(element)
		if err != nil {
			return err
		}
		elementString := *(*string)(unsafe.Pointer(&elementBytes))

		newElement := reflect.New(fieldValue.Type().Elem())
		err = genericSet(newElement.Interface(), elementString)
		if err != nil {
			return err
		}

		if newElement.Elem().Type().Kind() == reflect.String {
			correctedValue := newElement.Elem().String()
			correctedValue = correctedValue[1:len(correctedValue)-1]
			newElement.Elem().SetString(correctedValue)
		}

		fieldValue.Set(reflect.Append(fieldValue, newElement.Elem()))
	}
	return nil
}

func setMaps(fieldValue reflect.Value, stringValue string) error {
	ma := make(map[string]interface{})
	err := json.Unmarshal([]byte(stringValue[:]), &ma)
	if err != nil {
		return err
	}

	if fieldValue.IsNil() {
		fieldValue.Set(reflect.MakeMap(fieldValue.Type()))
	}

	for key, element := range ma {
		elementBytes, err := json.Marshal(&element)
		if err != nil {
			return err
		}
		elementString := *(*string)(unsafe.Pointer(&elementBytes))

		newElement := reflect.New(fieldValue.Type().Elem())
		err = genericSet(newElement.Interface(), elementString)
		if err != nil {
			return err
		}

		if newElement.Elem().Type().Kind() == reflect.String {
			correctedValue := newElement.Elem().String()
			correctedValue = correctedValue[1:len(correctedValue)-1]
			newElement.Elem().SetString(correctedValue)
		}

		fieldValue.SetMapIndex(reflect.ValueOf(key), newElement.Elem())
	}

	return nil
}