package easyconfig

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type LoadOption struct {
	EnvPrefix      string
	ConfigFilePath string // default is ./config/local.yaml
}

func LoadConfig(config interface{}, opt *LoadOption) error {
	var envPrefix string
	var configPath string
	if opt != nil {
		envPrefix = opt.EnvPrefix
		configPath = opt.ConfigFilePath
	}
	if configPath == "" {
		configPath = getConfigPath(envPrefix)
	}
	v := viper.New()
	if configPath != "" {
		configDir := filepath.Dir(configPath)
		configFile := strings.TrimSuffix(filepath.Base(configPath), ".yaml")
		v.SetConfigType("yaml")
		v.SetConfigName(configFile)
		v.AddConfigPath(configDir)
	}

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	defaultValues := ExtraDefaultValues(config)
	SetViperDefault(v, defaultValues)
	if configPath != "" {
		if err := v.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := LoadConfigFromViper(config, v); err != nil {
		return err
	}
	return nil
}

func SetViperDefault(v *viper.Viper, values map[string]interface{}) {
	for key, val := range values {
		v.SetDefault(key, val)
	}
}

func getConfigPath(envPrefix string) string {
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
	return path
}
