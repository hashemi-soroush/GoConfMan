package goconfman

import (
	"reflect"
)

func LoadFromMap(config interface{}, configMap map[string]interface{}) error {
	configValue := reflect.ValueOf(config)
	if configValue.Type().Kind() == reflect.Ptr {
		configValue = configValue.Elem()
	}

	//for key, element := range configMap {
	//	fieldValue := configValue.FieldByName(key)
	//	if fieldValue.Kind() == reflect.Invalid {
	//		return errors.New(fmt.Sprintf("%s isn't a field of %s", key, configValue.Type().PkgPath()))
	//	}
	//
	//	if fieldValue.Type().Kind() != reflect.Ptr {
	//		fieldValue = fieldValue.Addr()
	//	}
	//
	//	if fieldValue.Elem().Type().Kind() == reflect.Struct {
	//		elementMap, ok := element.(map[string]interface{})
	//		if ok == false {
	//			return errors.New(fmt.Sprintf("Can't convert this to map[string]interface{}: %v", element))
	//		}
	//
	//		err := LoadFromMap(fieldValue.Interface(), elementMap)
	//		if err != nil {
	//			return err
	//		}
	//	} else {
	//		elementBytes, err := json.Marshal(element)
	//		if err != nil {
	//			return err
	//		}
	//		elementString := *(*string)(unsafe.Pointer(&elementBytes))
	//
	//		err = json.Unmarshal([]byte(elementString[:]), fieldValue.Interface())
	//		//err = generic_set.GenericSet(fieldValue.Interface(), elementString)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//}
	return nil
}