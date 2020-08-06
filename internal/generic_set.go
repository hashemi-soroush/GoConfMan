package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func GenericSet(field interface{}, value string) {
	fieldValue := reflect.ValueOf(field)
	if fieldValue.Type().Kind() != reflect.Ptr {
		panic(errors.New("field must be a pointer"))
	}
	fieldType := fieldValue.Elem().Type()

	switch fieldType.Kind() {
	case reflect.String:
		setString(fieldValue, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		setInt(fieldValue, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		setUint(fieldValue, value)
	case reflect.Float32, reflect.Float64:
		setFloat(fieldValue, value)
	case reflect.Bool:
		setBool(fieldValue, value)
	case reflect.Array, reflect.Slice:
		setSlice(fieldValue, value)
	case reflect.Map:
		setMap(fieldValue, value)
	default:
		panic(errors.New(fmt.Sprintf("Type %s isn't supported by goconfman", fieldType.Kind().String())))
	}
}

func setBool(fieldValue reflect.Value, value string) {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	fieldValue.Elem().SetBool(boolValue)
}

func setInt(fieldValue reflect.Value, value string) {
	intValue, err := strconv.ParseInt(value, 0, fieldValue.Elem().Type().Bits())
	if err != nil {
		panic(err)
	}

	fieldValue.Elem().SetInt(intValue)
}

func setFloat(fieldValue reflect.Value, value string) {
	floatValue, err := strconv.ParseFloat(value, fieldValue.Elem().Type().Bits())
	if err != nil {
		panic(err)
	}

	fieldValue.Elem().SetFloat(floatValue)
}

func setUint(fieldValue reflect.Value, value string) {
	uintValue, err := strconv.ParseUint(value, 0, fieldValue.Elem().Type().Bits())
	if err != nil {
		panic(err)
	}

	fieldValue.Elem().SetUint(uintValue)
}

func setString(fieldValue reflect.Value, value string) {
	fieldValue.Elem().SetString(value)
}

func setMap(fieldValue reflect.Value, value string) {
	err := json.Unmarshal([]byte(value[:]), fieldValue.Interface())
	if err != nil {
		panic(err)
	}
}

func setSlice(fieldValue reflect.Value, value string) {
	err := json.Unmarshal([]byte(value), fieldValue.Interface())
	if err != nil {
		panic(err)
	}
}
