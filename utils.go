package easyconfig

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

var FinalKinds = map[reflect.Kind]bool{
	reflect.Bool:    true,
	reflect.Int:     true,
	reflect.Int8:    true,
	reflect.Int16:   true,
	reflect.Int32:   true,
	reflect.Int64:   true,
	reflect.Uint:    true,
	reflect.Uint8:   true,
	reflect.Uint16:  true,
	reflect.Uint32:  true,
	reflect.Uint64:  true,
	reflect.Float32: true,
	reflect.Float64: true,
	reflect.String:  true,

}

var NeedRecursiveKinds = map[reflect.Kind]bool{
	reflect.Interface: true,
	reflect.Ptr:       true,
	reflect.Uintptr:   true,
	reflect.Struct:    true,
}

var UnsupportKinds = map[reflect.Kind]bool{
	reflect.Complex64:     true,
	reflect.Complex128:    true,
	reflect.Chan:          true,
	reflect.Func:          true,
	reflect.UnsafePointer: true,
	reflect.Uintptr:       true,
	reflect.Slice: 		   true,
	reflect.Map: true,
}

func GetFieldConfigName(field reflect.StructField) string {
	configTagStr := field.Tag.Get("config")
	values := ConfigTagToValues(configTagStr)

	var name string
	if _, ok := values["name"]; ok {
		name = values["name"]
	} else {
		name = strcase.ToSnake(field.Name)
	}
	return name
}

func ExcludeFieldConfig(field reflect.StructField) bool {
	configTagStr := field.Tag.Get("config")
	return configTagStr == "-"
}

func GetViperKey(field reflect.StructField, prefix string) string {
	name := GetFieldConfigName(field)
	if prefix == "" {
		return name
	}
	return prefix + "." + name
}

func ConfigTagToValues(tagStr string) map[string]string {
	values := make(map[string]string)
	if tagStr == "" {
		return values
	}
	kvStrs := strings.Split(tagStr, ";")
	for _, kvStr := range kvStrs {
		if strings.Contains(kvStr, ":") {
			kvs := strings.Split(kvStr, ":")
			if len(kvs) >= 2 {
				k := strings.Trim(kvs[0], " ")
				v := kvs[1]
				values[k] = v
			}
		}
	}
	return values
}

func GetFieldConfigDefault(field reflect.StructField) (interface{}, error) {
	configTagStr := field.Tag.Get("config")
	values := ConfigTagToValues(configTagStr)
	var valueStr string
	var ok bool
	if valueStr, ok = values["default"]; !ok {
		return nil, nil
	}
	var value interface{}
	var err error
	if field.Type.Kind() == reflect.Ptr {
		value, err = StrConvertTo(valueStr, field.Type.Elem().Kind())
	} else {
		value, err = StrConvertTo(valueStr, field.Type.Kind())
	}
	if err != nil {
		fmt.Printf("field %v default value(%v) has wrong type, \n Error: %v \n", field.Name, valueStr, err)
		return nil, err
	}
	return value, nil
}

func AddPrefixToMapKey(prefix string, values map[string]interface{}) map[string]interface{} {
	if prefix == "" {
		return values
	}
	result := make(map[string]interface{})
	for k, v := range values {
		newKey := prefix + "." + k
		result[newKey] = v
	}
	return result
}

func StrConvertTo(s string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Float64:
		return cast.ToFloat64E(s)
	case reflect.Float32:
		return cast.ToFloat32E(s)
	case reflect.Int:
		return cast.ToIntE(s)
	case reflect.Int64:
		return cast.ToInt64E(s)
	case reflect.Int32:
		return cast.ToInt32E(s)
	case reflect.Int16:
		return cast.ToInt16E(s)
	case reflect.Int8:
		return cast.ToInt8E(s)
	case reflect.Uint:
		return cast.ToUintE(s)
	case reflect.Uint64:
		return cast.ToUint64E(s)
	case reflect.Uint32:
		return cast.ToUint32E(s)
	case reflect.Uint16:
		return cast.ToUint16E(s)
	case reflect.Uint8:
		return cast.ToUint8E(s)
	case reflect.String:
		return s, nil
	case reflect.Bool:
		if s == "true" {
			return true, nil
		} else {
			return false, nil
		}
	default:
		return nil, fmt.Errorf("unable to cast %#v to kind %T", s, kind)
	}
}

func UpdateMap(first, last map[string]interface{}) {
	if last == nil {
		return
	}
	for k, v := range last {
		first[k] = v
	}
}

func GetFieldViperKey(prefix, fieldKey string) string {
	if prefix == "" {
		return fieldKey
	}
	return strings.Join([]string{prefix, fieldKey}, ".")
}
