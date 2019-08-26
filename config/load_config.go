package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

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

var Config = new(OauthConfig)
var once sync.Once

type LoadOption struct {
	EnvPrefix         string
	DefaultConfigPath string // default is ./config/local.yaml
}

func getConfigPath(envPrefix, defaultPath string) string {
	var pathEnvVarName string
	if envPrefix == "" {
		pathEnvVarName = "CONFIG_PATH"
	} else {
		if strings.HasSuffix(envPrefix, "_") {
			pathEnvVarName = envPrefix + "CONFIG_PATH"
		} else {
			pathEnvVarName = envPrefix + "_" + "CONFIG_PATH"
		}
	}
	path := os.Getenv(pathEnvVarName)
	if path == "" {
		return defaultPath
	} else {
		return path
	}
}

//func GetDefaultValues(configPtr interface{}) map[string]interface{} {
//
//	var defaultValues = make(map[string]interface{})
//
//	v := reflect.ValueOf(configPtr)
//	k := v.Kind()
//	if k != reflect.Ptr {
//		panic("configPtr must a pointer of struct")
//	}
//	v.Float()
//
//	return defaultValues
//
//}

func GenYamlConfigTemplate() {

}
