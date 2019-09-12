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
	case reflect.Struct:
		val := reflect.New(field.Type)
		structVal := val.Elem()
		for i := 0; i < structVal.NumField(); i++ {
			_field := structVal.Type().Field(i)
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
			case reflect.Struct:
				val := reflect.New(field.Type.Elem())
				structVal := val.Elem()
				for i := 0; i < structVal.Type().NumField(); i++ {
					_field := structVal.Type().Field(i)
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

func SetFieldFromViper(field reflect.StructField, fieldVal reflect.Value, v *viper.Viper, prefix string) (err error) {
	viperKey := GetViperKey(field, prefix)
	switch fieldVal.Kind() {
	case reflect.Ptr:
		switch fieldVal.Elem().Kind() {
		case reflect.Int:
			val := v.GetInt(viperKey)
			fieldVal.Set(reflect.ValueOf(&val))
		case reflect.String:
			val := v.GetString(viperKey)
			fmt.Println("execute !")
			fieldVal.Set(reflect.ValueOf(&val))
		case reflect.Int64:
			val := v.GetInt64(viperKey)
			fieldVal.Set(reflect.ValueOf(&val))
		case reflect.Struct:
			return errors.New("unsupport field type: " + field.Type.Name())
		}
	case reflect.Int:
		fieldVal.SetInt(v.GetInt64(viperKey))
	case reflect.String:
		fieldVal.SetString(v.GetString(viperKey))
	case reflect.Bool:
		fieldVal.SetBool(v.GetBool(viperKey))
	case reflect.Uint:
		fieldVal.SetUint(v.GetUint64(viperKey))
	case reflect.Int64:
		fieldVal.SetInt(v.GetInt64(viperKey))
	case reflect.Struct:
		var nextPrefix string
		if prefix != "" {
			nextPrefix = prefix + "." + GetFieldConfigName(field)
		} else {
			nextPrefix = prefix
		}
		for i := 0; i < field.Type.NumField(); i++ {
			_field := field.Type.Field(i)
			_fieldVal := fieldVal.Field(i)
			SetFieldFromViper(_field, _fieldVal, v, nextPrefix)
		}
	default:
		return errors.New("unsupport field type: " + field.Type.Name())
	}
	return nil
}
