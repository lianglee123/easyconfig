# easyconfig
A package for load config from env variables and yaml file easily.

# usage
```go
package main

import (
    "fmt"
    "os"
	"github.com/lianglee123/easyconfig"
)

type MyConfig struct {
	Debug bool    `config:"default:true"`
	LogLevel string `config:"default:debug"`
	DB DBConfig
}

type DBConfig struct{
	Host string   `config:"127.0.0.1"`
	Port int   `config:"5432"`
	UserName string `config:"lianglee"`
	Pwd string  `config:"qwer1234"`
	DBName string  `config:"config_demo"`
}

func main() {
	opt := &easyconfig.LoadOption{
		EnvPrefix: "DEMO",
		ConfigFilePath: "./test.yaml",
	}
	config := &MyConfig{}
	os.Setenv("DEMO_LOG_LEVEL", "info")
	err := easyconfig.LoadConfig(config, opt)
	if err != nil {
		fmt.Printf("err happen when load config: %v \n", err)
		return
	}
	fmt.Printf("config: %+v", config)
}
```
config file `test.yaml`:
```yaml
db:
  pwd: abcdefg
```

execute result:
```
config: &{Debug:false LogLevel:info DB:{Host:127.0.0.1 Port:5432 UserName:lianglee Pwd:abcdefg DBName:config_demo}
```

# load priority
environment variables > yaml file > field tag(default value)

# next mission
- [ ] support more base type
- [ ] support slice
- [ ] use gomodule manage dependency