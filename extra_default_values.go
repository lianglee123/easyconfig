package easyconfig

import (
	"fmt"
	"reflect"
)

func ExtraDefaultValues(config interface{}) map[string]interface{} {
	cType := reflect.TypeOf(config)
	result := ExtraConfigValueFromType(cType, "")
	return result
}

func ExtraConfigValueFromType(_type reflect.Type, prefix string) map[string]interface{} {
	if _type.Kind() == reflect.Ptr {
		return ExtraConfigValueFromType(_type.Elem(), prefix)
	}
	var result = make(map[string]interface{})
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i)

		if ExcludeFieldConfig(field) {
			continue
		}

		fieldType := field.Type
		fieldKind := fieldType.Kind()
		if _, ok := UnsupportKinds[fieldKind]; ok {
			continue
		}

		if fieldKind == reflect.Struct || (fieldKind == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct) {

			_prefix := GetFieldConfigName(field)
			var nextPrefix string
			if prefix != "" {
				nextPrefix = prefix + "." + _prefix
			} else {
				nextPrefix = _prefix
			}
			values := ExtraConfigValueFromType(fieldType, nextPrefix)
			values = AddPrefixToMapKey(prefix, values)
			UpdateMap(result, values)
		}

		configKey := GetFieldConfigName(field)
		if prefix != "" {
			configKey = prefix + "." + configKey
		}
		defaultValue, err := GetFieldConfigDefault(field)
		if err != nil {
			fmt.Printf("get field %v default value error: %v", field.Name, err)
			continue
		}
		if defaultValue != nil {
			result[configKey] = defaultValue
		}

	}
	return result
}
