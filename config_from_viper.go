package easyconfig

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

func LoadConfigFromViper(config interface{}, v *viper.Viper) (err error) {
	configPtrVal := reflect.ValueOf(config)
	configPtrType := configPtrVal.Type()

	configStructType := configPtrType.Elem()
	configStructVal := configPtrVal.Elem()

	for i := 0; i < configStructType.NumField(); i++ {
		field := configStructType.Field(i)
		if ExcludeFieldConfig(field) {
			continue
		}
		fieldVal := configStructVal.Field(i)
		val, err := getValFromViper(v, "", field, fieldVal)
		if err != nil {
			continue
		}
		if fieldVal.CanSet() {
			fieldVal.Set(reflect.ValueOf(val))
		}
	}
	return nil
}

func getValFromViper(v *viper.Viper, prefix string, field reflect.StructField, fieldVal reflect.Value) (interface{}, error) {
	viperKey := GetViperKey(field, prefix)
	switch fieldVal.Kind() {
	case reflect.Int:
		return v.GetInt(viperKey), nil
	case reflect.String:
		return v.GetString(viperKey), nil
	case reflect.Bool:
		return v.GetBool(viperKey), nil
	case reflect.Int8:
		return int8(v.GetInt(viperKey)), nil
	case reflect.Int16:
		return int16(v.GetInt(viperKey)), nil
	case reflect.Int32:
		return int32(v.GetInt(viperKey)), nil
	case reflect.Int64:
		return int64(v.GetInt64(viperKey)), nil
	case reflect.Uint:
		return v.GetUint(viperKey), nil
	case reflect.Uint8:
		return uint8(v.GetUint(viperKey)), nil
	case reflect.Uint16:
		return uint16(v.GetUint(viperKey)), nil
	case reflect.Uint32:
		return uint32(v.GetUint(viperKey)), nil
	case reflect.Uint64:
		return uint64(v.GetUint(viperKey)), nil
	case reflect.Float32:
		return float32(v.GetFloat64(viperKey)), nil
	case reflect.Float64:
		return v.GetFloat64(viperKey), nil
	case reflect.Struct:
		val := reflect.New(field.Type)
		structVal := val.Elem()
		for i := 0; i < structVal.NumField(); i++ {
			_field := structVal.Type().Field(i)
			if ExcludeFieldConfig(_field) {
				continue
			}
			_fieldVal := structVal.Field(i)
			if _fieldVal.CanSet() {
				val, err := getValFromViper(v, viperKey, _field, _fieldVal)
				if err != nil {
					continue
				}
				_fieldVal.Set(reflect.ValueOf(val))
			}
		}
		return val.Elem().Interface(), nil
	case reflect.Ptr:
		{
			switch fieldVal.Type().Elem().Kind() {
			case reflect.String:
				val := v.GetString(viperKey)
				return &val, nil
			case reflect.Int:
				val := v.GetInt(viperKey)
				return &val, nil
			case reflect.Bool:
				val := v.GetBool(viperKey)
				return &val, nil
			case reflect.Int8:
				val := int8(v.GetInt(viperKey))
				return &val, nil
			case reflect.Int16:
				val := int16(v.GetInt(viperKey))
				return &val, nil
			case reflect.Int32:
				val := int32(v.GetInt(viperKey))
				return &val, nil
			case reflect.Int64:
				val := int64(v.GetInt64(viperKey))
				return &val, nil
			case reflect.Uint:
				val := v.GetUint(viperKey)
				return &val, nil
			case reflect.Uint8:
				val := uint8(v.GetUint(viperKey))
				return &val, nil
			case reflect.Uint16:
				val := uint16(v.GetUint(viperKey))
				return &val, nil
			case reflect.Uint32:
				val := uint32(v.GetUint(viperKey))
				return &val, nil
			case reflect.Uint64:
				val := uint64(v.GetUint(viperKey))
				return &val, nil
			case reflect.Float32:
				val := float32(v.GetFloat64(viperKey))
				return &val, nil
			case reflect.Float64:
				val := v.GetFloat64(viperKey)
				return &val, nil
			case reflect.Struct:
				val := reflect.New(field.Type.Elem())
				structVal := val.Elem()
				for i := 0; i < structVal.Type().NumField(); i++ {
					_field := structVal.Type().Field(i)
					if ExcludeFieldConfig(field) {
						continue
					}
					_fieldVal := structVal.Field(i)
					if _fieldVal.CanSet() {
						val, err := getValFromViper(v, viperKey, _field, _fieldVal)
						if err != nil {
							continue
						}
						_fieldVal.Set(reflect.ValueOf(val))
					}
				}
				return val.Interface(), nil
			default:
				errMsg := fmt.Sprintf("unsupport type of field(%v): %v", field.Name, fieldVal.Type().Kind().String())
				return nil, errors.New(errMsg)
			}

		}
	default:
		errMsg := fmt.Sprintf("unsupport type of field(%v): %v", field.Name, fieldVal.Type().Kind().String())
		return nil, errors.New(errMsg)
	}
	return nil, nil
}
