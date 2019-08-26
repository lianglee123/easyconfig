package config

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type Person struct {
	Name   string `config:"name:ming_zi;default:joy"`
	Age    int64  `config:"default:24"`
	Logger struct {
		Path         string `json:"path"`
		Level        string `json:"level"`
		Env          string `json:"env"`
		ReportCaller bool   `json:"reportcaller"`
	} `json:"logger"`
}

func TestReflect(t *testing.T) {
	var person interface{} = &Person{}
	ptrVal := reflect.ValueOf(person)
	fmt.Println("type: ", ptrVal.Type())
	fmt.Println("kind: ", ptrVal.Kind())

	structVal := ptrVal.Elem()
	structType := structVal.Type()

	defaultValues := map[string]string{}
	fmt.Println("structType PkgPath(): ", structType.PkgPath())
	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)

		fmt.Printf("structField: %+v\n", structField)
		fmt.Printf("structField.Name: %+v\n", structField.Name)
		fmt.Printf("structFieldTag: %+v\n", structField.Tag.Get("config"))

		configTag := structField.Tag.Get("config")
		tagValues := TagToValues(configTag)
		defaultValue, ok := tagValues["default"]
		name, ok := tagValues["name"]
		if !ok {
			name = strcase.ToSnake(structField.Name)
		}
		if ok {
			defaultValues[name] = defaultValue
		}

	}

}

func LoadConfig(config interface{}, opt *LoadOption) error {
	var envPrefix string
	var defaultConfigPath string
	if opt != nil {
		envPrefix = opt.EnvPrefix
		if opt.DefaultConfigPath != "" {
			defaultConfigPath = opt.DefaultConfigPath
		} else {
			defaultConfigPath = "./config/dev.yaml"
		}
	}
	configPath := getConfigPath(envPrefix, defaultConfigPath)
	configDir := filepath.Dir(configPath)
	configFile := strings.TrimSuffix(filepath.Base(configPath), ".yaml")

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.AddConfigPath(configDir)
	v.SetConfigFile(configFile)

	defaultValues := ExtraDefaultValues(config)
	SetViperDefault(defaultValues)
	SetValueToConfig(v, config)
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

var finalKinds = map[reflect.Kind]bool{
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

	reflect.Array: true,
	reflect.Slice: true,
}

var needRecursiveKinds = map[reflect.Kind]bool{
	reflect.Interface: true,
	reflect.Map:       true,
	reflect.Ptr:       true,
	reflect.Uintptr:   true,
	reflect.Struct:    true,
}

var unsurportKinds = map[reflect.Kind]bool{
	reflect.Complex64:     true,
	reflect.Complex128:    true,
	reflect.Chan:          true,
	reflect.Func:          true,
	reflect.UnsafePointer: true,
	reflect.Uintptr:   true,
}

// 如果没有明确设置默认，就设置为零值(指针除外)
func ExtraDefaultValues(config interface{}) map[string]interface{} {
	configPtrVal := reflect.ValueOf(config)
	configPtrType := configPtrVal.Type()
	configStructType := configPtrType.Elem()

	var result = make(map[string]string)

	for i := 0; i < configV; i++ {
		field := configStructType.Field(i)
		fieldVal := configPtrVal[1]
 		fieldType := field.Type
		fieldKind := fieldType.Kind()

		if _, ok := unsurportKinds[fieldKind]; ok {
			continue
		}

		var isPtr bool
		var elemType reflect.Type
		var elemVal reflect.Value
		if fieldKind == reflect.Ptr {
			elemVal = ()

		}
		if _, ok :=  finalKinds[fieldKind]; ok {
			configTagStr := field.Tag.Get("config")
			values := ConfigTagToValues(configTagStr)
		}


		if _, ok := needRecursiveKinds[fieldKind]; ok {
			continue
		}

		configTag := structField.Tag.Get("config")
		tagValues := TagToValues(configTag)
		defaultValue, ok := tagValues["default"]
		name, ok := tagValues["name"]
		if !ok {
			name = strcase.ToSnake(structField.Name)
		}
		if ok {
			defaultValues[name] = defaultValue
		}

	}

}

func extraDefaultValues(structPtr interface{}, prefix string) map[string]string {

}

func SetViperDefault(v *viper.Viper, values map[string]string) {
	for key, val := range values {
		v.SetDefault(key, val)
	}
}

func SetValueToConfig(v *viper.Viper, config interface{}) {
	v.AllSettings()
	v * viper.Vi
}

//
// tagStr format: key1:value1;key2:value2
func ConfigTagToValues(tagStr string) map[string]string {
	values := make(map[string]string)
	if tagStr == "" {
		return values
	}
	kvStrs := strings.Split(tagStr, ";")
	for _, kvStr := range kvStrs {
		if strings.Contains(kvStr, "=") {
			kvs := strings.Split(kvStr, "=")
			if len(kvs) >= 2 {
				k := strings.Trim(kvs[0], " ")
				v := kvs[1]
				values[k] = v
			}
		}
	}
	return values
}
