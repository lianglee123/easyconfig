package easyconfig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

type Person struct {
	Name   string `config:"name:ming_zi;default:joy"`
	Age    int64  `config:"default:24"`
	Logger struct {
		Path         string `config:"path"`
		Level        string `config:"level"`
		Env          string `config:"env"`
		ReportCaller bool   `config:"reportcaller"`
	} `config:"logger"`
}

func TestReflect(t *testing.T) {
	var person interface{} = &Person{}
	ptrVal := reflect.ValueOf(person)
	fmt.Println("type: ", ptrVal.Type())
	fmt.Println("kind: ", ptrVal.Kind())

	structVal := ptrVal.Elem()
	structType := structVal.Type()

	fmt.Println("structType PkgPath(): ", structType.PkgPath())
	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)

		fmt.Printf("structField: %+v\n", structField)
		fmt.Printf("structField.Name: %+v\n", structField.Name)
		fmt.Printf("structFieldTag: %+v\n", structField.Tag.Get("config"))
	}

}

type OauthConfig struct {
	Mode  string
	Redis struct {
		Host string `config:"name:host,default:localhost" `
		Port uint64 `config:"name:port;default:2370" `
	} `config:"redis"`
	OuterService struct {
		OauthAddr    string `config:"oauth_addr"`
		MetadataAddr string `config:"metadata_addr"`
	} `config:"outer_service;required"`

	EmailConfig struct {
		SMTPHost string
		SMTPPort int
		Pwd      string
		Account  string
	} `config:"name:email"`

	AccessTokenLifeTime  uint64
	RefreshTokenLifeTime uint64
	PosTokenLifetime     int64 `config:"name:config"`

	HTTPConfig struct {
		Host        string
		Port        string
		TemplateDir string
		StaticDir   string
	}

	RPCConfig struct {
		Host string
		Port string
	}

	DBConfig map[string]*DataBaseConfig `config:"name:db_config"`

	LoggerConfig struct {
		Level  string
		Format string
	} `config:"name:logger_config"`
}

type DataBaseConfig struct {
	Driver   string `config:"driver"`
	Host     string `config:"host"`
	Port     string `config:"port"`
	Name     string
	User     string
	Pwd      string
	PoolSize string
}

// test final type
func TestLoadConfig_1(t *testing.T) {
	type Person struct {
		Name string `config:"default:joy"`
		Age  int    `config:"default:24"`
	}

	p := &Person{}
	opt := &LoadOption{
		EnvPrefix:      "TEST",
		ConfigFilePath: "./test.yaml",
	}
	err := LoadConfig(p, opt)
	assert.Equal(t, err, nil)
	assert.Equal(t, p.Name, "joy.lee")
	assert.Equal(t, p.Age, 28)

	opt = &LoadOption{
		EnvPrefix:      "TEST",
		ConfigFilePath: "",
	}

	err = LoadConfig(p, opt)
	assert.NoError(t, err)
	assert.Equal(t, p.Name, "joy")
	assert.Equal(t, p.Age, 24)

	os.Setenv("TEST_NAME", "joyddd")
	os.Setenv("TEST_AGE", "30")
	err = LoadConfig(p, opt)
	assert.NoError(t, err)
	assert.Equal(t, p.Name, "joyddd")
	assert.Equal(t, p.Age, 30)
}

func TestLoadConfig_2(t *testing.T) {
	type Son struct {
		Name string `config:"default:joy2.0"`
		Age  int    `config:"default:1"`
	}
	type Person struct {
		Aag  int    `config:"default:24"`
		Name string `config:"default:joy"`
		Son  Son
	}

	p := &Person{}
	opt := &LoadOption{
		EnvPrefix:      "TEST",
		ConfigFilePath: "./test.yaml",
	}
	err := LoadConfig(p, opt)
	assert.Equal(t, err, nil)
}
